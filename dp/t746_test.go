package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func TestT746(t *testing.T) {
	assert.Equal(t, 15, minCostClimbingStairs([]int{10, 15, 20}))
	assert.Equal(t, 6, minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
