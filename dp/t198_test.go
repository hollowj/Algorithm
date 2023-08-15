package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func rob(nums []int) int {
//		dp := make([][]int, len(nums))
//		for i := 0; i < len(nums); i++ {
//			dp[i] = make([]int, 2)
//		}
//		dp[0][1] = nums[0]
//		for i := 1; i < len(nums); i++ {
//			dp[i][0] = max(dp[i-1][1], dp[i-1][0])
//			dp[i][1] = dp[i-1][0] + nums[i]
//		}
//		return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
//	}
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(nums)-1]
}
func TestT198(t *testing.T) {
	assert.Equal(t, 4, rob([]int{1, 2, 3, 1}))
	assert.Equal(t, 12, rob([]int{2, 7, 9, 3, 1}))
}
