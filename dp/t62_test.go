package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
func TestT62(t *testing.T) {
	assert.Equal(t, 28, uniquePaths(3, 7))
	assert.Equal(t, 3, uniquePaths(3, 2))
	assert.Equal(t, 28, uniquePaths(7, 3))
	assert.Equal(t, 6, uniquePaths(3, 3))
	assert.Equal(t, 1, uniquePaths(1, 1))
}
