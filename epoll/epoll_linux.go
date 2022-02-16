//go:build linux
// +build linux

/**
 * @Time: 2022/2/16 21:44
 * @Author: yt.yin
 */

package epoll

import (
	"golang.org/x/sys/unix"
	"net"
	"reflect"
	"sync"
	"syscall"
)

// 定义多路复用的结构体
type epoll struct {
	fd   int
	cons map[int]net.Conn
	lock *sync.RWMutex
}

// NewEpoll 初始化一个多路复用的对象
func NewEpoll() (*epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &epoll{
		fd:   fd,
		lock: &sync.RWMutex{},
		cons: make(map[int]net.Conn),
	}, nil
}

// Add 新增连接
func (e *epoll) Add(conn net.Conn) error {
	// 获取与连接关联的文件描述符
	fd := PrySysfd(conn)
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
func (e *epoll) Remove(conn net.Conn) error {
	// 移除连接前先关闭防止客户端如果是单片机会持有死链接
	_ = conn.Close()
	fd := PrySysfd(conn)
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

// Wait 等待
func (e *epoll) Wait() ([]net.Conn, error) {
	events := make([]unix.EpollEvent, 100)
retry:
	n, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
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
		conns = append(conns, conn)
	}
	return conns, nil
}

// PrySysfd 通过反射手段获取系统文件描述符
// FD 是一个文件描述符。 net 和 os 包使用这种类型作为表示网络连接或操作系统文件的较大类型的字段。
// Sysfd 系统文件描述符。 在关闭之前不可变。
// 通过反射获取 系统文件描述符
func PrySysfd(conn net.Conn) int {
	// 通过反射获取不可见属性 conn 见 net/tcpsock.go:86
	tcpC := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	// 通过反射获取网络文件描述符fd 见 net/net.go:171
	fdV := tcpC.FieldByName("fd")
	// 通过反射获取 poll.FD 的值
	pfdV := reflect.Indirect(fdV).FieldByName("pfd")
	// 通过反射获取 系统文件描述符
	return int(pfdV.FieldByName("Sysfd").Int())
}
