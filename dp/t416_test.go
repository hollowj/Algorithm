package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	result := sum / 2
	dp := make([]int, result+1)
	for i := 0; i < len(nums); i++ {
		for j := result; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}

	}
	return dp[result] == result
}

func TestT416(t *testing.T) {
	assert.Equal(t, true, canPartition([]int{1, 5, 11, 5}))
	assert.Equal(t, false, canPartition([]int{1, 2, 3, 5}))
}
