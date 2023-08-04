package tanxin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func wiggleMaxLength(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 1
	delta := 0
	for i := 1; i < len(nums); i++ {
		tmp := nums[i] - nums[i-1]
		if tmp == 0 {
			continue
		}
		if delta == 0 || delta*tmp < 0 {
			count++
			delta = tmp
		}
	}
	return count
}

func TestT376(t *testing.T) {
	assert.Equal(t, 1, wiggleMaxLength([]int{1}))
	assert.Equal(t, 2, wiggleMaxLength([]int{1, 2}))
	assert.Equal(t, 6, wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	assert.Equal(t, 7, wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
	assert.Equal(t, 2, wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	assert.Equal(t, 1, wiggleMaxLength([]int{0, 0}))
}
