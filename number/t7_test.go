package number

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func reverse(x int) int {
	target := 0
	for x != 0 {
		target *= 10
		target += x % 10
		x = x / 10
	}
	if target > math.MaxInt32 || target < math.MinInt32 {
		return 0
	}
	return target
}

func TestT7(t *testing.T) {
	assert.Equal(t, reverse(123), 321)
	assert.Equal(t, reverse(-123), -321)
	assert.Equal(t, reverse(120), 21)
	assert.Equal(t, reverse(0), 0)
}
