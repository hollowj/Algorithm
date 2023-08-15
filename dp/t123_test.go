package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit123(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}
	//1，第一次持有股票
	//2，第一次不持有股票
	//3，第二次持有股票
	//4，第二次不持有股票
	dp[0][1] = -prices[0]
	dp[0][3] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][1] = max(dp[i-1][1], -prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[len(prices)-1][4]
}
func TestT123(t *testing.T) {
	assert.Equal(t, 6, maxProfit123([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	assert.Equal(t, 4, maxProfit123([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, maxProfit123([]int{7, 6, 4, 3, 1}))
	assert.Equal(t, 0, maxProfit123([]int{1}))
}
