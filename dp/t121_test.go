package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][1] = max(dp[i-1][1], -prices[i])
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	}
	return dp[len(prices)-1][0]
}

//	func maxProfit(prices []int) int {
//		iMin := prices[0]
//		delta := 0
//		for i := 1; i < len(prices); i++ {
//			price := prices[i]
//			if price < iMin {
//				iMin = price
//			} else {
//				if price-iMin > delta {
//					delta = price - iMin
//				}
//			}
//		}
//		return delta
//	}
func TestT121(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}
