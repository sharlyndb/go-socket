/**
 * @Time: 2022/2/18 12:09
 * @Author: yt.yin
 */

package encoding

type Encoder interface {
	// Encode 编码
	Encode([]byte) ([]byte, error)

	// Decode 解码
	Decode([]byte) (data []byte, err error)
}
