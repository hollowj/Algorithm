package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[int32]int)
	for _, ch := range magazine {
		m[ch] += 1
	}
	for _, ch := range ransomNote {
		m[ch]--
		if m[ch] < 0 {
			return false
		}
	}
	return true
}

func TestT383(t *testing.T) {
	assert.Equal(t, false, canConstruct("a", "b"))
	assert.Equal(t, false, canConstruct("aa", "ab"))
	assert.Equal(t, true, canConstruct("aa", "aab"))
}
