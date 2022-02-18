/**
 * @Time: 2022/2/17 20:48
 * @Author: yt.yin
 */

package encoding

import (
	"bytes"
	"fmt"
)

type BinaryEncoder struct{}

// ByteToBinStr 将byte 以8个bit位的形式展示
func(e *BinaryEncoder) ByteToBinStr(b byte) string {
	return fmt.Sprintf("%08b", b)
}

// BytesToBinStr 将byte数组转成8个bit位一组的字符串
func (e *BinaryEncoder) BytesToBinStr(bs []byte) string {
	if len(bs) <= 0 {
		return ""
	}
	buf := bytes.NewBuffer([]byte{})
	for _, v := range bs {
		buf.WriteString(fmt.Sprintf("%08b", v))
	}
	return buf.String()
}

// BytesToBinStrWithSplit 将byte数组转8个bit一组的字符串并且带分割符
func (e *BinaryEncoder) BytesToBinStrWithSplit(bs []byte,split string) string {
	length := len(bs)
	if length <= 0 {
		return ""
	}
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < length-1; i++ {
		v := bs[i]
		buf.WriteString(fmt.Sprintf("%08b", v))
		buf.WriteString(split)
	}
	buf.WriteString(fmt.Sprintf("%08b",bs[length-1]))
	return buf.String()
}