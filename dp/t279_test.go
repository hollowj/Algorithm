package dp

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	m := int(math.Floor(math.Sqrt(float64(n))))
	for i := 1; i <= m; i++ {
		num := int(math.Pow(float64(i), float64(2)))
		for j := num; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-num]+1)
		}
	}
	return dp[n]
}

func TestT279(t *testing.T) {
	assert.Equal(t, 3, numSquares(12))
	assert.Equal(t, 2, numSquares(13))
}
