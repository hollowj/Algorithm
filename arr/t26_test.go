package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func removeDuplicates(nums []int) int {
//		iRealIndex := 0
//		iIndex := iRealIndex + 1
//		for iIndex < len(nums) {
//			for iIndex < len(nums) && nums[iIndex] == nums[iRealIndex] {
//				iIndex++
//			}
//			if iRealIndex < len(nums) && iIndex < len(nums) {
//				iRealIndex++
//				nums[iRealIndex] = nums[iIndex]
//			}
//
//		}
//		return iRealIndex + 1
//	}
func removeDuplicates(nums []int) int {
	iSlow := 0
	iFast := iSlow + 1
	for iFast < len(nums) {
		if nums[iFast] != nums[iSlow] {
			iSlow++
			nums[iSlow] = nums[iFast]
		}
		iFast++
	}
	return iSlow + 1
}
func TestT26(t *testing.T) {
	assert.Equal(t, 2, removeDuplicates([]int{1, 1, 2}))
	assert.Equal(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
