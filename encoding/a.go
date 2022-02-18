/**
 * @Time: 2022/2/17 18:02
 * @Author: yt.yin
 */

package encoding

// 定义包内的一个解码组，涵盖定义的常见的解码器
type encodeGroup struct{
	// ASCII 编码器
	AsciiEncoder
	// BCD 编码器
	BcdEncoder
	// 二进制编码器
	BinaryEncoder
	// HEX 编码器
	HexEncoder
}

// EncodeGroup 对外统一开放一个解码组
var EncodeGroup = new(encodeGroup)
