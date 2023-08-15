package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit188(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 1+2*k)
	}

	for i := 0; i < k; i++ {
		dp[0][2*i+1] = -prices[0]
	}
	for i := 1; i < len(prices); i++ {
		for j := 1; j < 2*k+1; j++ {
			if j%2 == 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			}
		}
	}
	return dp[len(prices)-1][2*k]
}
func TestT188(t *testing.T) {
	assert.Equal(t, 2, maxProfit188(2, []int{2, 4, 1}))
	assert.Equal(t, 7, maxProfit188(2, []int{3, 2, 6, 5, 0, 3}))
}
