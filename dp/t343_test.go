package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, (i-j)*j))
		}
	}
	return dp[n]
}
func TestT343(t *testing.T) {
	assert.Equal(t, 1, integerBreak(2))
	assert.Equal(t, 36, integerBreak(10))
	assert.Equal(t, 18, integerBreak(8))
}
