/**
 * @Time: 2022/2/17 11:59
 * @Author: yt.yin
 */

package iface

// MsgI 通用的message接口
type MsgI interface {

	// MsgID 消息ID
	MsgID()   uint32

	// DataLen 获取消息数据段长度
	DataLen() uint32

	// Data 获取消息内容
	Data()    []byte

	// SetMsgID 设置消息ID
	SetMsgID()

	// SetDataLen 设置消息数据段长度
	SetDataLen()

	// SetData 设置消息内容
	SetData()
}

// TLVMsgI 定义几个常用的消息模版，行业常用的有 TLV，8583等
type TLVMsgI interface {

	MsgI

	// Tag 消息标识
	Tag()     []byte

	// Begin 消息头或开始，常用的有 FFFF,55AA 等
	Begin()   []byte

	// End 消息结束符,常用的有 FF,回车符等
	End()     []byte

	// ICV 完整性校验值，常用的有BCC校验 CRC校验
	ICV()     []byte

	// SetTag 设置消息标识
	SetTag()

	// SetBegin 设置消息头或开始，常用的有 FFFF,55AA 等
	SetBegin()

	// SetEnd 设置消息结束符,常用的有 FF,回车符等
	SetEnd()

	// SetICV 设置 完整性校验值
	SetICV()
}