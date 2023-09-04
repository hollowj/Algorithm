package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
func TestT344(t *testing.T) {
	bytes := []byte("olleh")
	reverseString(bytes)
	assert.Equal(t, []byte("olleh"), bytes)
}
