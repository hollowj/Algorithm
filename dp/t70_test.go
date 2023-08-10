package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestT70(t *testing.T) {
	assert.Equal(t, 2, climbStairs(2))
	assert.Equal(t, 3, climbStairs(3))
	assert.Equal(t, 5, climbStairs(4))
}
