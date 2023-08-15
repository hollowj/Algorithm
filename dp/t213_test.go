package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func rob213(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}
func robRange(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	dp := make([]int, len(nums))
	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])
	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[end]
}
func TestT213(t *testing.T) {
	assert.Equal(t, 3, rob213([]int{2, 3, 2}))
	assert.Equal(t, 4, rob213([]int{1, 2, 3, 1}))
	assert.Equal(t, 3, rob213([]int{1, 2, 3}))
	assert.Equal(t, 0, rob213([]int{0, 0}))
}
