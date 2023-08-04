package tanxin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func maxSubArray(nums []int) int {
//	sum := math.MinInt
//	s := 0
//	for i := 0; i < len(nums); i++ {
//		s += nums[i]
//		if s > sum {
//			sum = s
//		}
//		if s < 0 {
//			s = 0
//		}
//	}
//	return sum
//}

//	func maxSubArray(nums []int) int {
//		sum := -math.MaxInt
//		for i := 0; i < len(nums); i++ {
//			for j := i; j < len(nums); j++ {
//				_sum := 0
//				for k := i; k < j; k++ {
//					_sum += nums[k]
//				}
//				sum = int(math.Max(float64(sum), float64(_sum)))
//			}
//		}
//		return sum
//	}

//func maxSubArray(nums []int) int {
//	dp := make([][]int, len(nums))
//	for i := 0; i < len(nums); i++ {
//		ints := make([]int, 0)
//		for j := 0; j < len(nums); j++ {
//			ints = append(ints, math.MinInt)
//		}
//		dp[i] = ints
//	}
//	sum := math.MinInt
//	for i := 0; i < len(nums); i++ {
//		for j := i; j < len(nums); j++ {
//			if i == j {
//				dp[i][j] = nums[i]
//			} else {
//				dp[i][j] = dp[i][j-1] + nums[j]
//			}
//			if dp[i][j] > sum {
//				sum = dp[i][j]
//			}
//		}
//	}
//	return sum
//}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	sum := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > sum {
			sum = dp[i]
		}
	}
	return sum
}
func TestT53(t *testing.T) {
	//assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	//assert.Equal(t, 1, maxSubArray([]int{1}))
	//assert.Equal(t, 23, maxSubArray([]int{5, 4, -1, 7, 8}))
	assert.Equal(t, -1, maxSubArray([]int{-1, -2}))

}
