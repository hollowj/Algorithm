package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func search(nums []int, target int) int {
//		return searchByMid(nums, target, 0, len(nums)-1)
//	}
//
//	func searchByMid(nums []int, target, begin, end int) int {
//		if end < begin {
//			return -1
//		}
//		iMid := begin + (end-begin)/2
//		if nums[iMid] > target {
//			return searchByMid(nums, target, begin, iMid-1)
//		} else if nums[iMid] == target {
//			return iMid
//		} else {
//			return searchByMid(nums, target, iMid+1, end)
//		}
//	}
//
//	func search(nums []int, target int) int {
//		iBegin := 0
//		iEnd := len(nums)
//		for iBegin < iEnd {
//			iMid := iBegin + (iEnd-iBegin)/2
//			if nums[iMid] > target {
//				iEnd = iMid - 1
//			} else if nums[iMid] < target {
//				iBegin = iMid + 1
//			} else {
//				return iMid
//			}
//		}
//		if iBegin >= len(nums) || iEnd < 0 {
//			return -1
//		}
//		if nums[iBegin] == target {
//			return iBegin
//		}
//		return -1
//	}
//
//	func search(nums []int, target int) int {
//		iBegin := 0
//		iEnd := len(nums) - 1
//		for iBegin <= iEnd {
//			iMid := iBegin + (iEnd-iBegin)/2
//			if nums[iMid] > target {
//				iEnd = iMid - 1
//			} else if nums[iMid] < target {
//				iBegin = iMid + 1
//			} else {
//				return iMid
//			}
//		}
//
//		return -1
//	}
func search(nums []int, target int) int {
	iBegin := 0
	iEnd := len(nums)
	for iBegin < iEnd {
		iMid := iBegin + (iEnd-iBegin)/2
		if nums[iMid] > target {
			iEnd = iMid
		} else if nums[iMid] < target {
			iBegin = iMid + 1
		} else {
			return iMid
		}
	}

	return -1
}
func TestT704(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 13))
}
