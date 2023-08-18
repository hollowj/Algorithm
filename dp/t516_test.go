package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	longestLen := 0
	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
		longestLen = 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
			if dp[i][j] > longestLen {
				longestLen = dp[i][j]
			}
		}
	}
	return longestLen
}

func TestT516(t *testing.T) {
	assert.Equal(t, 4, longestPalindromeSubseq("bbbab"))
	assert.Equal(t, 2, longestPalindromeSubseq("cbbd"))
	assert.Equal(t, 1, longestPalindromeSubseq("a"))
}
