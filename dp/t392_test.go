package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isSubsequence(s string, t string) bool {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(t)+1; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(s)][len(t)] == len(s)
}
func TestName(t *testing.T) {
	assert.Equal(t, true, isSubsequence("abc", "ahbgdc"))
	assert.Equal(t, false, isSubsequence("axc", "ahbgdc"))
}
