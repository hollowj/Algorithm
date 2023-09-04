package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
*
i i+2-j-1
0 1
1 1
*/
//func reverseStr(s string, k int) string {
//	bytes := make([]byte, len(s))
//	reverseBegin := 0
//	if len(s)-reverseBegin < k {
//		k = len(s) - reverseBegin
//	}
//	j := 0
//	for i := 0; i < len(s); i++ {
//		char := s[i]
//		if reverseBegin != -1 {
//			bytes[i] = s[reverseBegin+k-1-j]
//		} else {
//			bytes[i] = char
//		}
//		j++
//		if j == k {
//			j = 0
//			if reverseBegin == -1 {
//				reverseBegin = i + 1
//				if len(s)-reverseBegin < k {
//					k = len(s) - reverseBegin
//				}
//			} else {
//				reverseBegin = -1
//			}
//		}
//	}
//
//	return string(bytes)
//}
func reverseString_1(s []byte, begin, end int) {
	for i := begin; i < begin+(end-begin)/2; i++ {
		s[i], s[begin+end-1-i] = s[begin+end-1-i], s[i]
	}
}
func reverseStr(s string, k int) string {
	bytes := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		begin := i
		end := begin + k - 1
		if begin+k > len(s) {
			end = len(s) - 1
		}
		for begin < end {
			bytes[begin], bytes[end] = bytes[end], bytes[begin]
			begin++
			end--
		}
	}
	return string(bytes)
}

func TestT541(t *testing.T) {
	assert.Equal(t, "bacdfeg", reverseStr("abcdefg", 2))
	assert.Equal(t, "bacd", reverseStr("abcd", 2))
	assert.Equal(t, "a", reverseStr("a", 2))
	assert.Equal(t, "cbadefg", reverseStr("abcdefg", 3))
	assert.Equal(t, "cbad", reverseStr("abcd", 3))
	assert.Equal(t, "dcba", reverseStr("abcd", 4))
}
