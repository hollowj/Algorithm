package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minSubArrayLen(target int, nums []int) int {
	iSlow := 0
	iSum := 0
	iMinLength := 0
	for iFast := 0; iFast < len(nums); iFast++ {
		iSum += nums[iFast]
		for iSum >= target {
			size := iFast - iSlow + 1
			if iMinLength == 0 || size < iMinLength {
				iMinLength = size
			}
			iSum -= nums[iSlow]
			iSlow++
		}
	}

	return iMinLength
}
func TestT209(t *testing.T) {
	assert.Equal(t, 2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	assert.Equal(t, 1, minSubArrayLen(4, []int{1, 4, 4}))
	assert.Equal(t, 0, minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
