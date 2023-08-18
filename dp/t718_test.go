package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2))
	}
	maxLen := 0
	for i := 0; i < len(nums1); i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
			maxLen = 1
		}
	}
	for i := 0; i < len(nums2); i++ {
		if nums1[0] == nums2[i] {
			dp[0][i] = 1
			maxLen = 1
		}
	}
	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}

			if dp[i][j] > maxLen {
				maxLen = dp[i][j]
			}
		}
	}
	return maxLen
}

//	func findLength(nums1 []int, nums2 []int) int {
//		dp := make([][]int, len(nums1)+1)
//		for i := 0; i <= len(nums1); i++ {
//			dp[i] = make([]int, len(nums2)+1)
//		}
//		maxLen := 0
//		for i := 1; i <= len(nums1); i++ {
//			for j := 1; j <= len(nums2); j++ {
//				if nums1[i-1] == nums2[j-1] {
//					dp[i][j] = dp[i-1][j-1] + 1
//				}
//
//				if dp[i][j] > maxLen {
//					maxLen = dp[i][j]
//				}
//			}
//		}
//		return maxLen
//	}
func TestT718(t *testing.T) {
	assert.Equal(t, 3, findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	assert.Equal(t, 5, findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))
}
