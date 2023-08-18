package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	maxLength := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxLength {
			maxLength = dp[i]
		}
	}
	return maxLength
}

func TestT300(t *testing.T) {
	assert.Equal(t, 4, lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	assert.Equal(t, 4, lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	assert.Equal(t, 1, lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
