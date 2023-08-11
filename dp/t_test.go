package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// index 1 物品重量
//
//	func maxPrice(weight, price []int, bag int) int {
//		dp := make([][]int, len(weight))
//		for i := 0; i < len(weight); i++ {
//			dp[i] = make([]int, bag+1)
//		}
//		for i := weight[0]; i <= bag; i++ {
//			dp[0][i] = price[0]
//		}
//		for i := 1; i < len(weight); i++ {
//			for j := weight[0]; j <= bag; j++ {
//				if j < weight[i] {
//					dp[i][j] = dp[i-1][j]
//				} else {
//					dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+price[i])
//				}
//			}
//		}
//		return dp[len(weight)-1][bag]
//	}

func maxPrice(weight, price []int, bag int) int {
	dp := make([]int, bag+1)
	for i := 0; i < len(weight); i++ {
		for j := bag; j >= weight[i]; j-- {
			dp[j] = max(dp[j], dp[j-weight[i]]+price[i])
		}
	}
	return dp[bag]
}

func TestT(t *testing.T) {
	assert.Equal(t, 35, maxPrice([]int{1, 3, 4}, []int{15, 20, 30}, 4))
}
