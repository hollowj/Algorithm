package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 0; i < len(s)+1; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
func TestT115(t *testing.T) {
	assert.Equal(t, 3, numDistinct("rabbbit", "rabbit"))
	assert.Equal(t, 5, numDistinct("babgbag", "bag"))
}
