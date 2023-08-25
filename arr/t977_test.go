package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func sortedSquares(nums []int) []int {
//		for i := 0; i < len(nums); i++ {
//			nums[i] = nums[i] * nums[i]
//		}
//		sort.Ints(nums)
//		return nums
//	}
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	iLeft := 0
	iRight := len(nums) - 1
	iResultIndex := len(nums) - 1
	for iLeft <= iRight {
		a := nums[iLeft] * nums[iLeft]
		b := nums[iRight] * nums[iRight]
		if a > b {
			iLeft++
			result[iResultIndex] = a
		} else {
			iRight--
			result[iResultIndex] = b
		}
		iResultIndex--
	}
	return result
}
func TestT977(t *testing.T) {
	assert.Equal(t, []int{0, 1, 9, 16, 100}, sortedSquares([]int{-4, -1, 0, 3, 10}))
	assert.Equal(t, []int{4, 9, 9, 49, 121}, sortedSquares([]int{-7, -3, 2, 3, 11}))
}
