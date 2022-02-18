/**
 * @Time: 2022/2/17 18:02
 * @Author: yt.yin
 */

package encoding

// 定义包内的一个解码组，涵盖定义的常见的解码器
type encodeGroup struct{
	AsciiEncoder
	BcdEncoder
	BinaryEncoder
	HexEncoder
}

// EncodeGroup 对外统一开放一个解码组
var EncodeGroup = new(encodeGroup)
