/**
 * @Time: 2022/2/19 12:23
 * @Author: yt.yin
 */

package manager

import (
	"errors"
	"net"
	"reflect"
	"sync"
	"syscall"

	"github.com/goworkeryyt/go-socket/server/connect"
	"golang.org/x/sys/unix"
)

// EpollManager 基于 linux epoll 实现连接管理
type EpollManager struct{
	fd   int
	cons map[int]connect.ConnI
	lock *sync.RWMutex
}

// New 初始化连接管理池
func New() (*EpollManager, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &EpollManager{
		fd:   fd,
		lock: &sync.RWMutex{},
		cons: make(map[int]connect.ConnI),
	}, nil
}

// Add 添加连接到连接池
func (e *EpollManager) Add(conn connect.ConnI) error {
	if conn == nil || conn.Conn() == nil{
		return errors.New("连接为空")
	}
	// 获取与连接关联的文件描述符
	fd := tcpConnFd(conn.Conn())
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.cons[fd] = conn
	return nil
}

// Remove 移除连接
func (e *EpollManager) Remove(conn connect.ConnI) error{
	if conn == nil || conn.Conn() == nil {
		return errors.New("连接为空")
	}
	// 移除连接前先关闭防止客户端如果是单片机会持有死链接
	_ = conn.Conn().Close()
	fd := tcpConnFd(conn.Conn())
retry:
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		if err == unix.EINTR {
			goto retry
		}
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.cons, fd)
	return nil
}

// Conn 根据连接的id从连接池里获取连接
func (e *EpollManager) Conn(connID int) (connect.ConnI, error) {
	connI := e.cons[connID]
	if connI != nil {
		return connI,nil
	}
	return nil,errors.New("连接不存在")
}

// Size 获取连接数量
func (e *EpollManager) Size() int {
	return len(e.cons)
}

// Vacuum 清空连接池
func (e *EpollManager) Vacuum() {
	// TODO
}


// Wait 等待 maxEvent 为每次最多读取的事件数量,默认配置 100,最大1000
func (e *EpollManager) Wait(maxEvent int) ([]net.Conn, error) {
	if maxEvent < 100 {
		maxEvent = 100
	}else if maxEvent > 1000 {
		maxEvent = 1000
	}
	events := make([]unix.EpollEvent, maxEvent)
retry:
	n, err := unix.EpollWait(e.fd, events, maxEvent)
	if err != nil {
		// 判断是否有 EINTR错误
		if err == unix.EINTR {
			goto retry
		}
		return nil, err
	}
	e.lock.RLock()
	defer e.lock.RUnlock()
	var conns []net.Conn
	for i := 0; i < n; i++ {
		conn := e.cons[int(events[i].Fd)]
		conns = append(conns, conn.Conn())
	}
	return conns, nil
}

// 获取连接的 fd
// FD 是一个文件描述符。 net 和 os 包使用这种类型作为表示网络连接或操作系统文件的较大类型的字段。
// Sysfd 系统文件描述符。 在关闭之前不可变。
// 通过反射获取 系统文件描述符
func tcpConnFd(conn net.Conn) int {
	// 通过反射获取不可见属性 conn 见 net/tcpsock.go:86
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	// 通过反射获取网络文件描述符fd 见 net/net.go:171
	fdV := tcpConn.FieldByName("fd")
	// 通过反射获取 poll.FD 的值
	pfdV := reflect.Indirect(fdV).FieldByName("pfd")
	// 通过反射获取 系统文件描述符
	return int(pfdV.FieldByName("Sysfd").Int())
}


