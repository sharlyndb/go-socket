/**
 * @Time: 2022/2/17 18:50
 * @Author: yt.yin
 */

package encoding

import (
	"bytes"
	"encoding/hex"
	"errors"
	"strconv"
)

// BcdEncoder bcd 编码一般只针对数字
type BcdEncoder struct{}

// Uint64ToBcd 64位无符号正数转bcd
func(b *BcdEncoder) Uint64ToBcd(v uint64) ([]byte,error) {
	numStr := strconv.FormatUint(v, 10)
	if len(numStr) % 2 != 0 {
		// 左边补一个0
		var sb bytes.Buffer
		sb.WriteString("0")
		sb.WriteString(numStr)
		numStr = sb.String()
	}
	return bcd([]byte(numStr))
}

// UintToBcd 无符号正数转bcd
func(b *BcdEncoder) UintToBcd(v uint) ([]byte,error) {
	numStr := strconv.Itoa(int(v))
	if len(numStr) % 2 != 0 {
		var sb bytes.Buffer
		sb.WriteString("0")
		sb.WriteString(numStr)
		numStr = sb.String()
	}
	return bcd([]byte(numStr))
}

// IntToBcd int 转bcd
func (b *BcdEncoder) IntToBcd(v int) ([]byte,error) {
	if v < 0 {
		return nil,errors.New("负数不能转BCD格式")
	}
	numStr := strconv.Itoa(v)
	if len(numStr) % 2 != 0 {
		var sb bytes.Buffer
		sb.WriteString("0")
		sb.WriteString(numStr)
		numStr = sb.String()
	}
	return bcd([]byte(numStr))
}

// BcdToInt bcd 转int
func (b *BcdEncoder) BcdToInt(data []byte) (int,error){
	hexStr := hex.EncodeToString(data)
	parseUint, err := strconv.ParseUint(hexStr, 10, 32)
	return int(parseUint), err
}

// BcdToUint32 bcd转 unit32
func (b *BcdEncoder) BcdToUint32(data []byte) (uint32,error) {
	hexStr := hex.EncodeToString(data)
	parseUint, err := strconv.ParseUint(hexStr, 10, 32)
	return uint32(parseUint), err
}

// BcdToUint64 bcd转 unit64
func (b *BcdEncoder) BcdToUint64(data []byte) (uint64,error){
	hexStr := hex.EncodeToString(data)
	return strconv.ParseUint(hexStr, 10, 64)
}

// 将 ascii 中的数字编码为 bcd（确保 len(data) % 2 == 0）
func bcd(data []byte) ([]byte,error) {
	out := make([]byte, len(data)/2+1)
	n, err := hex.Decode(out, data)
	if err != nil {
		return nil, err
	}
	return out[:n],nil
}