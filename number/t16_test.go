package number

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	num := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		iLeft, iRight := i+1, len(nums)-1
		for iLeft < iRight {
			x := nums[i] + nums[iLeft] + nums[iRight]
			if math.Abs(float64(target-x)) < math.Abs(float64(target-num)) {
				num = x
			}
			if x > target {
				iRight--
			} else if x < target {
				iLeft++
			} else {
				return target
			}
		}
	}

	return num
}
func TestT16(t *testing.T) {
	assert.Equal(t, 2, threeSumClosest([]int{-1, 2, 1, -4}, 1))
	assert.Equal(t, 0, threeSumClosest([]int{0, 0, 0}, 1))
}
