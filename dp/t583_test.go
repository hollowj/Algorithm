package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func minDistance(word1 string, word2 string) int {
//		dp := make([][]int, len(word1)+1)
//		for i := 0; i < len(word1)+1; i++ {
//			dp[i] = make([]int, len(word2)+1)
//		}
//		for i := 1; i < len(word1)+1; i++ {
//			for j := 1; j < len(word2)+1; j++ {
//				if word1[i-1] == word2[j-1] {
//					dp[i][j] = dp[i-1][j-1] + 1
//				} else {
//					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//				}
//			}
//		}
//		return len(word1) + len(word2) - dp[len(word1)][len(word2)]*2
//	}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(word1)+1; i++ {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i < len(word1)+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(word2)+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < len(word1)+1; i++ {
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func TestT583(t *testing.T) {
	assert.Equal(t, 2, minDistance("sea", "eat"))
	assert.Equal(t, 4, minDistance("leetcode", "etco"))
	assert.Equal(t, 2, minDistance("a", "b"))
}
