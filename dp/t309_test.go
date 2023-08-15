package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit309(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 4)
	}
	//0 持有
	//1 卖出
	//2 冷冻
	dp[0][0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = dp[i-1][1]
	}
	return dp[len(prices)-1][1]
}

func TestT309(t *testing.T) {
	assert.Equal(t, 3, maxProfit309([]int{1, 2, 3, 0, 2}))
	assert.Equal(t, 0, maxProfit309([]int{1}))
}
