/**
 * @Time: 2022/2/18 02:45
 * @Author: yt.yin
 */

package encoding

import (
	"encoding/hex"
	"strings"
)

// HexEncoder hex 编码器
type HexEncoder struct{}

// Encode 编码 []byte{0x55, 0xAA} 被转成 55AA
func (h *HexEncoder) Encode (src []byte)  (dst []byte, err error){
	s := strings.ToUpper(hex.EncodeToString(src))
	return []byte(s),nil
}

// Decode 解码 AABBCC 转成字节数组 []byte{0xAA, 0xBB, 0xCC}
func (h *HexEncoder) Decode(src []byte) (dst []byte, err error){
	return hex.DecodeString(string(src))
}

// BytesToHex 字节数组转hex
// []byte{0x55, 0xAA} 被转成 55AA
func (e *HexEncoder) BytesToHex(data []byte) string{
	return strings.ToUpper(hex.EncodeToString(data))
}

// HexToBytes 将hex 字符串转成 byte数组
// AABBCC 转成字节数组 []byte{0xAA, 0xBB, 0xCC}
func (e *HexEncoder) HexToBytes(hexStr string) ([]byte, error) {
	return hex.DecodeString(hexStr)
}

// HexBCC 计算BCC校验码
func (e *HexEncoder) HexBCC(hexStr string) string {
	hexToBytes, err := e.HexToBytes(hexStr)
	if err != nil {
		return ""
	}
	length := len(hexToBytes)
	if length < 1 {
		return ""
	}
	bcc := hexToBytes[0]
	if length > 1 {
		for i := 1; i < length; i++ {
			bcc = bcc ^ hexToBytes[i]
		}
	}
	return e.BytesToHex([]byte{bcc & 0xFF})
}

// BytesBCC 计算 BCC
func BytesBCC(bytes []byte) byte {
	bcc := bytes[0]
	if len(bytes) > 1 {
		for i := 1; i < len(bytes); i++ {
			bcc = bcc ^ bytes[i]
		}
	}
	return bcc & 0xFF
}