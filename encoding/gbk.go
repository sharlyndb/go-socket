/**
 * @Time: 2022/2/18 15:11
 * @Author: yt.yin
 */

package encoding

import "golang.org/x/text/encoding/simplifiedchinese"

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	GBK     = Charset("GBK")
)

// GBKEncoder GBK 编码器, go 语言默认使用UTF-8编码,国内很多协议使用的编码规范为 GBK,或 GB18030
type GBKEncoder struct{}

// Encode 编码 utf-8 转到 gbk
func (g *GBKEncoder) Encode(src []byte) ([]byte, error) {
	return simplifiedchinese.GBK.NewEncoder().Bytes(src)
}

// Decode 解码 gbk 转到 utf-8
func (g *GBKEncoder) Decode(src []byte) (data []byte, err error) {
	return simplifiedchinese.GBK.NewEncoder().Bytes(src)
}

// UTF8StrToGBK UTF 字符串转 GBK
func UTF8StrToGBK(str string) (string,error) {
	decodeBytes,err :=simplifiedchinese.GBK.NewEncoder().Bytes([]byte(str))
	if err != nil {
		return "", err
	}
	return string(decodeBytes),nil
}

// GBKStrToUTF8 GBK 字符串转UTF
func GBKStrToUTF8(str string) (string,error) {
	decodeBytes,err :=simplifiedchinese.GBK.NewDecoder().Bytes([]byte(str))
	if err != nil {
		return "", err
	}
	return string(decodeBytes),nil
}

// UTF8StrToGB18030 UTF 字符串转 GB18030
func UTF8StrToGB18030(str string) (string,error) {
	decodeBytes,err :=simplifiedchinese.GB18030.NewEncoder().Bytes([]byte(str))
	if err != nil {
		return "", err
	}
	return string(decodeBytes),nil
}

// GB18030StrToUTF8 GBK 字符串转UTF
func GB18030StrToUTF8(str string) (string,error) {
	decodeBytes,err :=simplifiedchinese.GB18030.NewDecoder().Bytes([]byte(str))
	if err != nil {
		return "", err
	}
	return string(decodeBytes),nil
}



