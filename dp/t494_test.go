package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func findTargetSumWays(nums []int, target int) int {
//	sum := 0
//	for i := 0; i < len(nums); i++ {
//		sum += nums[i]
//	}
//	if (target+sum)%2 == 1 {
//		return 0
//	}
//	if abs(target) > sum {
//		return 0
//	}
//	result := (target + sum) / 2
//	dp := make([]int, result+1)
//	dp[0] = 1
//	for i := 0; i < len(nums); i++ {
//		for j := result; j >= nums[i]; j-- {
//			dp[j] += dp[j-nums[i]]
//		}
//	}
//	return dp[result]
//}

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if (target+sum)%2 == 1 {
		return 0
	}
	if abs(target) > sum {
		return 0
	}
	result := (target + sum) / 2
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, result+1)
	}
	dp[0][0] = 1
	if result >= nums[0] {
		dp[0][nums[0]] += 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j <= result; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j-nums[i]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(nums)-1][result]
}

func TestT494(t *testing.T) {
	assert.Equal(t, 5, findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	assert.Equal(t, 1, findTargetSumWays([]int{1}, 1))
}
