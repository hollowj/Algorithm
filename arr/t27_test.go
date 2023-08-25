package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func removeElement(nums []int, val int) int {
//		size := len(nums)
//		for i := 0; i < size; i++ {
//			if nums[i] == val {
//				size--
//				for j := i; j < size; j++ {
//					nums[j] = nums[j+1]
//				}
//				i--
//			}
//
//		}
//		return size
//	}

//	func removeElement(nums []int, val int) int {
//		iSlowIndex := 0
//		for i := 0; i < len(nums); i++ {
//			if nums[i] != val {
//				nums[iSlowIndex] = nums[i]
//				iSlowIndex++
//			}
//		}
//		return iSlowIndex
//	}
//
//	func removeElement(nums []int, val int) int {
//		if len(nums) == 0 {
//			return 0
//		}
//		iSlowIndex := 0
//		iFastIndex := len(nums) - 1
//		for iSlowIndex <= iFastIndex {
//			for iSlowIndex <= iFastIndex && nums[iSlowIndex] != val {
//				iSlowIndex++
//			}
//			for iSlowIndex <= iFastIndex && nums[iFastIndex] == val {
//				iFastIndex--
//
//			}
//			if iSlowIndex < iFastIndex {
//				nums[iSlowIndex] = nums[iFastIndex]
//				iFastIndex--
//				iSlowIndex++
//			}
//		}
//		return iSlowIndex
//	}
func removeElement(nums []int, val int) int {
	iSlow := 0
	for iFast := 0; iFast < len(nums); iFast++ {
		if nums[iFast] != val {
			nums[iSlow] = nums[iFast]
			iSlow++
		}
	}
	return iSlow
}
func TestT27(t *testing.T) {
	assert.Equal(t, 2, removeElement([]int{3, 2, 2, 3}, 3))
	assert.Equal(t, 5, removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	assert.Equal(t, 0, removeElement([]int{}, 0))
	assert.Equal(t, 0, removeElement([]int{1}, 1))
}
