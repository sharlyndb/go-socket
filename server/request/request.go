/**
 * @Time: 2022/2/17 15:19
 * @Author: yt.yin
 */

package request

import "github.com/goworkeryyt/go-socket/server/connect"

// RequestI 客户端请求的连接信息和数据
type RequestI interface{

	// Conn 客户端请求的连接信息
	Conn() connect.ConnI

	// Data 获取当前请求的数据
	Data() []byte

	// MsgID 获取请求的消息ID
	MsgID() uint32
}
