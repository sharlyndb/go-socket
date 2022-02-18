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