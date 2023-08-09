package tanxin

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func monotoneIncreasingDigits(n int) int {
//		ints := make([]int, 0)
//		for n > 0 {
//			remain := n % 10
//			ints = append(ints, remain)
//			n = n / 10
//		}
//		result := 0
//		bFix := false
//		for i := len(ints) - 1; i >= 0; i-- {
//			if bFix {
//				result += 9 * int(math.Pow10(i))
//				continue
//			}
//			if i == 0 || ints[i] <= ints[i-1] {
//				result += ints[i] * int(math.Pow10(i))
//			} else {
//				result += (ints[i] - 1) * int(math.Pow10(i))
//				bFix = true
//			}
//		}
//		if bFix {
//			return monotoneIncreasingDigits(result)
//		}
//		return result
//	}
func monotoneIncreasingDigits(n int) int {
	nStr := fmt.Sprintf("%d", n)
	int32s := make([]int, len(nStr))
	for i, ch := range nStr {
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			panic(err)
		}
		int32s[i] = num
	}
	result := 0
	fixPos := 0
	for i := len(int32s) - 1; i > 0; i-- {
		if int32s[i-1] > int32s[i] {
			int32s[i-1] = int32s[i-1] - 1
			fixPos = i
		}
	}
	for i := 0; i < len(int32s); i++ {
		if fixPos == 0 || i < fixPos {
			result += int32s[i] * int(math.Pow10(len(int32s)-i-1))
		} else {
			result += 9 * int(math.Pow10(len(int32s)-i-1))
		}
	}

	return result
}
func TestT738(t *testing.T) {
	assert.Equal(t, 9, monotoneIncreasingDigits(10))
	assert.Equal(t, 1234, monotoneIncreasingDigits(1234))
	assert.Equal(t, 299, monotoneIncreasingDigits(332))
}
