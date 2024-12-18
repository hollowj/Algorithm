package string

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func myAtoi(s string) int {
	for i, ch := range s {
		if ch != ' ' {
			if i > 0 {
				s = s[i:]
			}
			break
		}
	}
	target := 0
	flag := 1
	for i, ch := range s {
		if i == 0 && ch == '-' {
			flag = -1
			continue
		}
		if i == 0 && ch == '+' {
			continue
		}
		if ch >= '0' && ch <= '9' {
			target *= 10
			i := ch - '0'
			target += int(i)
			if flag < 0 && target*-1 < math.MinInt32 {
				return math.MinInt32
			}
			if flag > 0 && target > math.MaxInt32 {
				return math.MaxInt32
			}
		} else {
			break
		}

	}
	return target * flag

}

func TestT08(t *testing.T) {
	fmt.Println(math.MaxInt64)
	assert.Equal(t, 42, myAtoi("42"))
	assert.Equal(t, -42, myAtoi(" -042"))
	assert.Equal(t, 1337, myAtoi("1337c0d3"))
	assert.Equal(t, 0, myAtoi("0-1"))
	assert.Equal(t, 0, myAtoi("words and 987"))
	assert.Equal(t, 0, myAtoi("   +0 123"))
	assert.Equal(t, 2147483647, myAtoi("9223372036854775808"))
	assert.Equal(t, -2147483648, myAtoi("-2147483648"))
}
