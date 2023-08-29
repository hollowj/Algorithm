package hash

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func threeSum(nums []int) [][]int {
//		sort.Ints(nums)
//		result := make([][]int, 0)
//		m := make(map[int]int)
//		for i, num := range nums {
//			for j := i + 1; j < len(nums); j++ {
//				t := 0 - num - nums[j]
//				if pos, ok := m[t]; ok && i != pos && j != pos {
//					result = append(result, []int{num, nums[j], t})
//				}
//			}
//		}
//		return result
//	}
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
				left++
			} else {
				left++
			}
		}
	}

	return result
}
func TestT15(t *testing.T) {
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.Equal(t, [][]int{}, threeSum([]int{0, 1, 1}))
	assert.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0}))
	assert.Equal(t, [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}
