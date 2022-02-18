/**
 * @Time: 2022/2/18 03:09
 * @Author: yt.yin
 */

package simple

// SimpleMsgI 通用的message接口
type SimpleMsgI interface {

	// MsgID 消息ID
	MsgID()      uint32

	// DataLen 获取消息数据段长度
	DataLen()    int

	// GetData 获取消息内容
	GetData()    []byte

	// SetMsgID 设置消息ID
	SetMsgID(id uint32)

	// SetDataLen 设置消息数据段长度
	SetDataLen(length int)

	// SetData 设置消息内容
	SetData(data []byte)
}

// SimpleMsg 消息
type SimpleMsg struct {

	//消息的ID
	ID      uint32

	//消息的长度
	Len     int

	//消息的内容
	Data    []byte
}

// NewSimpleMsg new 一个message
func NewSimpleMsg(ID uint32, data []byte) *SimpleMsg {
	return &SimpleMsg{
		Len:     len(data),
		ID:      ID,
		Data:    data,
	}
}

// MsgID 获取消息ID
func (m *SimpleMsg) MsgID() uint32 {
	return m.ID
}

// DataLen 获取消息长度
func (m *SimpleMsg) DataLen() int {
	return m.Len
}

// GetData 获取消息内容
func (m *SimpleMsg) GetData() []byte {
	return m.Data
}

// SetMsgID 设置消息ID
func (m *SimpleMsg) SetMsgID(id uint32) {
	m.ID = id
}

// SetDataLen 设置消息长度
func (m *SimpleMsg) SetDataLen(length int) {
	m.Len = length
}

// SetData 设置数据
func (m *SimpleMsg) SetData(data []byte) {
	m.Data = data
}
