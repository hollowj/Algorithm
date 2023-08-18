package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s)+1)
	}
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					count++
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					count++
					dp[i][j] = true
				}
			}
		}
	}
	return count
}
func TestT647(t *testing.T) {
	assert.Equal(t, 3, countSubstrings("abc"))
	assert.Equal(t, 6, countSubstrings("aaa"))
}
