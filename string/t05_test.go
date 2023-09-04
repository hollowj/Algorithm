package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func replaceSpace(s string) string {
	bytes := make([]byte, 0, len(s))
	for _, ch := range s {
		if ch == ' ' {
			bytes = append(bytes, byte('%'), byte('2'), byte('0'))
		} else {
			bytes = append(bytes, byte(ch))
		}
	}
	return string(bytes)
}

func TestT05(t *testing.T) {
	assert.Equal(t, "We%20are%20happy.", replaceSpace("We are happy."))
}
