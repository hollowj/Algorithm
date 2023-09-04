package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func reverseLeftWords(s string, n int) string {
//		bytes := []byte(s)
//		tmp := make([]byte, 0, n)
//		slow := 0
//		for i := 0; i < len(s); i++ {
//			if i < n {
//				tmp = append(tmp, bytes[i])
//			} else {
//				bytes[slow] = bytes[i]
//				slow++
//			}
//		}
//		for i := 0; i < len(tmp); i++ {
//			bytes[slow] = tmp[i]
//			slow++
//		}
//		return string(bytes)
//	}
func reverseLeftWords(s string, n int) string {
	bytes := []byte(s)
	reverse(bytes[:n])
	reverse(bytes[n:])
	reverse(bytes)
	return string(bytes)
}

func TestT58(t *testing.T) {
	assert.Equal(t, "cdefgab", reverseLeftWords("abcdefg", 2))
	assert.Equal(t, "umghlrlose", reverseLeftWords("lrloseumgh", 6))
}
