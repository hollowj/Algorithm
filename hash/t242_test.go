package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isAnagram(s string, t string) bool {
	m := make(map[int32]int)
	for _, ch := range s {
		count := m[ch]
		m[ch] = count + 1
	}
	for _, ch := range t {
		count, ok := m[ch]
		if !ok {
			return false
		}
		count--
		if count == 0 {
			delete(m, ch)
		} else {
			m[ch] = count
		}
	}
	return len(m) == 0
}

//	func isAnagram(s string, t string) bool {
//		m := make([]int, 26)
//		for _, ch := range s {
//			m[ch-97]++
//		}
//		for _, ch := range t {
//			m[ch-97]--
//		}
//		for i := 0; i < 26; i++ {
//			if m[i] != 0 {
//				return false
//			}
//		}
//		return true
//	}
func TestT242(t *testing.T) {
	assert.Equal(t, true, isAnagram("anagram", "nagaram"))
	assert.Equal(t, false, isAnagram("rat", "car"))
}
