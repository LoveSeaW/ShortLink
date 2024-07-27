package base62

import (
	"math"
	"strings"
)

// const base62Str = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
// const base62Str = `J0rs12O5TUV8IW7D9aBdXeCfghiMQj3klmop6qtuvbcwx4zAEFGHKLNnPRYSZy`

var (
	baseStr    string
	baseStrLen uint64
)

func MustInit(bs string) {
	if len(bs) == 0 {
		panic("need base string")
	}
	baseStr = bs
	baseStrLen = uint64(len(bs))
}

// 转换为62进制
func Int62ToString(seq uint64) string {
	if seq == 0 {
		return string(baseStr[0])
	}
	bl := []byte{}
	for seq > 0 {
		mod := seq % 62
		div := seq / 62
		bl = append(bl, baseStr[mod])
		seq = div
	}
	return string(reverse(bl))
}

func StringToIn62(s string) (seq uint64) {
	bl := reverse([]byte(s))
	for index, b := range bl {
		base := math.Pow(62, float64(index))
		seq = uint64(strings.Index(baseStr, string(b))) * uint64(base)
	}
	return seq
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
