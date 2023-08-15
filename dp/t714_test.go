package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit714(prices []int, fee int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[len(prices)-1][0]
}
func TestT714(t *testing.T) {
	assert.Equal(t, 8, maxProfit714([]int{1, 3, 2, 8, 4, 9}, 2))
	assert.Equal(t, 6, maxProfit714([]int{1, 3, 7, 5, 10, 3}, 3))
}
