package string

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverseWords(s string) string {
	bytes := []byte(s)
	slow := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			bytes[slow] = s[i]
			slow++
		} else {
			if i > 0 && s[i-1] != ' ' {
				bytes[slow] = s[i]
				slow++
			}
		}
	}
	if bytes[slow-1] == ' ' {
		slow--
	}
	bytes = bytes[:slow]
	reverse(bytes)
	begin := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			reverse(bytes[begin:i])
			begin = i + 1
		}
		if i == len(bytes)-1 {
			reverse(bytes[begin : i+1])
		}
	}
	return string(bytes)
}
func reverse(bytes []byte) {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}
}

func TestT151(t *testing.T) {
	assert.Equal(t, "blue is sky the", reverseWords("the sky is blue"))
	assert.Equal(t, "world hello", reverseWords("  hello world  "))
	assert.Equal(t, "example good a", reverseWords("a good   example"))
	n := "hwj"
	fmt.Println(n[:1])
}
