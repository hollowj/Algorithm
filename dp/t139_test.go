package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func wordBreak(s string, wordDict []string) bool {
//		dp := make([]bool, len(s)+1)
//		dp[0] = true
//		for i := 1; i <= len(s); i++ {
//			for j := 0; j < len(wordDict); j++ {
//				charLen := len(wordDict[j])
//				if i >= charLen && dp[i-charLen] && s[i-charLen:i] == wordDict[j] {
//					dp[i] = true
//				}
//			}
//		}
//		return dp[len(s)]
//	}
func wordBreak(s string, wordDict []string) bool {
	set := make(map[string]bool)
	for _, ch := range wordDict {
		set[ch] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j <= i; j++ {
			subStr := s[j:i]
			if dp[j] && set[subStr] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}
func TestT139(t *testing.T) {
	assert.Equal(t, true, wordBreak("leetcode", []string{"leet", "code"}))
	assert.Equal(t, true, wordBreak("applepenapple", []string{"apple", "pen"}))
	assert.Equal(t, false, wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}
