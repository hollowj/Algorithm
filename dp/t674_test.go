package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLength := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < nums[i] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}
	return maxLength
}

func TestT674(t *testing.T) {
	assert.Equal(t, 3, findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	assert.Equal(t, 1, findLengthOfLCIS([]int{2, 2, 2, 2, 2}))
}
