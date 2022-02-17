/**
 * @Time: 2022/2/17 11:41
 * @Author: yt.yin
 */

package iface

import (
	"context"
	"net"
)

// ConnI 定义连接接口
type ConnI interface {

	// Start 启动连接
	Start()

	// Stop 停止连接
	Stop()

	// Context 返回上下文，获取连接推出状态等
	Context() context.Context

	// TCPConn 获取原始的 tcp socket 连接
	TCPConn() *net.TCPConn

	// ConnID 连接ID
	ConnID() int

	// RemoteAddr 获取远程客户端地址信息
	RemoteAddr() net.Addr

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
