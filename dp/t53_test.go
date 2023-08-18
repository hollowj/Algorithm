package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] >= 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}

func TestT53(t *testing.T) {
	assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 1, maxSubArray([]int{1}))
	assert.Equal(t, 23, maxSubArray([]int{5, 4, -1, 7, 8}))
}
