package number

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// timeout
//
//	func divide(dividend int, divisor int) int {
//		dividend32 := int32(dividend)
//		divisor32 := int32(divisor)
//		sign := 1
//		if (dividend32^divisor32)>>31&0x1 == 1 {
//			sign = -1
//		}
//		count := 0
//		j := math.Abs(float64(dividend32))
//		k := math.Abs(float64(divisor32))
//		for j > 0 {
//			j -= k
//			if j >= 0 {
//				count++
//			}
//		}
//		count = count * sign
//		if count > math.MaxInt32 {
//			return math.MaxInt32
//		}
//		if count < math.MinInt32 {
//			return math.MinInt32
//		}
//		return count
//	}
func absInt32(num int32) int32 {
	if num < 0 {
		return -1 * num
	}
	return num
}
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 {
		if divisor == 1 {
			return math.MinInt32
		} else if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == 0 {
		return 0
	}
	dividend32 := int32(dividend)
	divisor32 := int32(divisor)
	reverse := false
	if dividend32 > 0 {
		dividend32 = -dividend32
		reverse = !reverse
	}
	if divisor32 > 0 {
		divisor32 = -divisor32
		reverse = !reverse
	}
	count := 0
	for dividend32 <= divisor32 {
		tmpCount := 1
		tmpDivisor := divisor32
		for dividend32 <= tmpDivisor && tmpDivisor < 0 {
			count += tmpCount
			dividend32 -= tmpDivisor
			tmpDivisor = tmpDivisor << 1
			tmpCount = tmpCount << 1
		}
	}
	if reverse {
		count = -count
	}
	if count > math.MaxInt32 {
		return math.MaxInt32
	}
	if count < math.MinInt32 {
		return math.MinInt32
	}
	return count
}
func TestT29(t *testing.T) {
	assert.Equal(t, 3, divide(10, 3))
	assert.Equal(t, -2, divide(7, -3))
	assert.Equal(t, 1, divide(1, 1))
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
	assert.Equal(t, 1073741823, divide(2147483647, 2))
	assert.Equal(t, -1073741824, divide(-2147483648, 2))
	assert.Equal(t, -1, divide(1100540749, -1090366779))
}
