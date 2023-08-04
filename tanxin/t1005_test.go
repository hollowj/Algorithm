package tanxin

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func largestSumAfterKNegations(nums []int, k int) int {
//	sort.Slice(nums, func(i, j int) bool {
//		return nums[i] < nums[j]
//	})
//	for i := 0; i < len(nums) && k > 0; i++ {
//		if nums[i] < 0 {
//			nums[i] = -nums[i]
//			k--
//		} else {
//			break
//		}
//	}
//	iSum := 0
//	iMin := nums[0]
//	for _, num := range nums {
//		iSum += num
//		if num < iMin {
//			iMin = num
//		}
//	}
//	if k > 0 && k%2 != 0 {
//		iSum = iSum - iMin - iMin
//	}
//	return iSum
//}

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k%2 == 1 {
		nums[len(nums)-1] *= -1
	}
	iSum := 0
	for _, num := range nums {
		iSum += num
	}
	return iSum
}
func TestT1005(t *testing.T) {
	assert.Equal(t, 5, largestSumAfterKNegations([]int{4, 2, 3}, 1))
	assert.Equal(t, 6, largestSumAfterKNegations([]int{3, -1, 0, 2}, 3))
	assert.Equal(t, 13, largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2))
	assert.Equal(t, 22, largestSumAfterKNegations([]int{-8, 3, -5, -3, -5, -2}, 6))
}
