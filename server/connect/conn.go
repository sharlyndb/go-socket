/**
 * @Time: 2022/2/17 11:41
 * @Author: yt.yin
 */

package connect

import (
	"context"
	"net"
	"sync"

	"github.com/goworkeryyt/go-socket/server/handler"
)

// ConnI 定义连接接口
type ConnI interface {

	// Start 启动连接
	Start()

	// Stop 停止连接
	Stop()

	// Context 返回上下文，获取连接推出状态等
	Context() context.Context

	// Conn 获取原始的 socket 连接
	Conn() net.Conn

	// ConnID 连接ID
	ConnID() int

	// RemoteAddr 获取远程客户端地址信息
	RemoteAddr() net.Addr

	// SetMsgMsgHandler 设置消息处理器
	SetMsgMsgHandler()

	// SendMsg 发送消息给TCP客户端,无缓存
	SendMsg(msgId uint32, data []byte) error

	// SendBufMsg 发送消息给TCP客户端(有缓冲)
	SendBufMsg(msgID uint32, data []byte) error

	// SetAttr 设置链接属性
	SetAttr(key string, value interface{})

	// Attr 获取链接属性
	Attr(key string) (interface{}, error)

	// DelAttr 移除链接属性
	DelAttr(key string)
}

// SimpleConn 简单连接实现
type SimpleConn struct {

	// 当前的socket连接
	conn net.Conn

	// 消息处理器
	MsgHandler handler.MsgHandlerI

	// 连接的上下文
	ctx    context.Context

	// 通知连接退出停止
	cancel context.CancelFunc

	// 无缓冲管道
	noBufferChan chan []byte

	// 有缓存管道
	BufferChan   chan []byte

	// 读写锁
	sync.RWMutex

	// 连接属性
	attrs        map[string]interface{}

	// 读写属性锁
	attrLock     sync.Mutex

	// 当前连接的状态，closed 为true 连接已经处于关闭状态
	isClosed     bool

}

// Start 启动连接让连接开始工作
func (s *SimpleConn) Start() {
	panic("implement me")
}

// Stop 停止连接
func (s *SimpleConn) Stop() {
	s.cancel()
}

func (s *SimpleConn) Context() context.Context {
	panic("implement me")
}

func (s *SimpleConn) Conn() net.Conn {
	panic("implement me")
}

func (s *SimpleConn) ConnID() int {
	panic("implement me")
}

func (s *SimpleConn) RemoteAddr() net.Addr {
	panic("implement me")
}

func (s *SimpleConn) SetMsgMsgHandler() {
	panic("implement me")
}

func (s *SimpleConn) SendMsg(msgId uint32, data []byte) error {
	panic("implement me")
}

func (s *SimpleConn) SendBufMsg(msgID uint32, data []byte) error {
	panic("implement me")
}

func (s *SimpleConn) SetAttr(key string, value interface{}) {
	panic("implement me")
}

func (s *SimpleConn) Attr(key string) (interface{}, error) {
	panic("implement me")
}

func (s *SimpleConn) DelAttr(key string) {
	panic("implement me")
}



