package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func TestT96(t *testing.T) {
	assert.Equal(t, 5, numTrees(3))
	assert.Equal(t, 1, numTrees(1))
}
