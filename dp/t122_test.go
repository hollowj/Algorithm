package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit122(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
func TestT122(t *testing.T) {
	assert.Equal(t, 7, maxProfit122([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 4, maxProfit122([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, maxProfit122([]int{7, 6, 4, 3, 1}))
}
