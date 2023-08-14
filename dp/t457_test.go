package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	//dp[i][j]表示i个0，j个1的最大子集长度
	for _, str := range strs {
		count0 := 0
		count1 := 0
		for _, ch := range str {
			if ch == '0' {
				count0++
			} else {
				count1++
			}
		}
		for i := m; i >= count0; i-- {
			for j := n; j >= count1; j-- {
				dp[i][j] = max(dp[i][j], dp[i-count0][j-count1]+1)
			}
		}
	}
	return dp[m][n]
}

func TestT457(t *testing.T) {
	assert.Equal(t, 4, findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	assert.Equal(t, 2, findMaxForm([]string{"10", "0", "1"}, 1, 1))
}
