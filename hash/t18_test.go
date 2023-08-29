package hash

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] > target && (nums[i] > 0 || target > 0) {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j] > target && (nums[i]+nums[j] > 0 || target > 0) {
				break
			}
			left := j + 1
			right := len(nums) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
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
	}

	return result
}
func TestT18(t *testing.T) {
	//assert.Equal(t, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	//assert.Equal(t, [][]int{{2, 2, 2, 2}}, fourSum([]int{2, 2, 2, 2, 2}, 8))
	assert.Equal(t, [][]int{{-5, -4, -3, 1}}, fourSum([]int{1, -2, -5, -4, -3, 3, 3, 5}, -11))
}
