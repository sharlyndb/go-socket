/**
 * @Time: 2022/2/17 18:01
 * @Author: yt.yin
 */

package encoding

import (
	"fmt"
)

type AsciiEncoder struct{}

// Encode ASCII 编码
func (a AsciiEncoder) Encode(data []byte) ([]byte, error) {
	bs := make([]byte,len(data))
	for _, d := range data {
		if d > 127 {
			return nil, fmt.Errorf("无效的 ASCII 字符：'%s'", string(d))
		}
		bs = append(bs,d)
	}
	return bs, nil
}

// Decode 指定长度 ASCII 解码
func (a AsciiEncoder) Decode(data []byte) ([]byte, error) {
	bs := make([]byte,len(data))
	for _, d := range data {
		if d > 127 {
			return nil, fmt.Errorf("无效的 ASCII 字符：'%s'", string(d))
		}
		bs = append(bs,d)
	}
	return bs, nil
}

// AssignLenDecode 指定长度 ASCII 解码
func (a AsciiEncoder) AssignLenDecode(data []byte, length int) ([]byte, int, error) {
	// 要解码的长度必须不小于原始字节数组的长度
	if len(data) < length {
		return nil, 0, fmt.Errorf("要解码的长度必须不小于原始字节数组的长度， 期望长度 %d， 实际长度 %d", length, len(data))
	}
	data = data[:length]
	var out []byte
	for _, d := range data {
		if d > 127 {
			return nil, 0, fmt.Errorf("无效的 ASCII 字符：'%s'", string(d))
		}
		out = append(out, d)
	}
	return out, length, nil
}

