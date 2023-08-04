package tanxin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func jump(nums []int) int {
	lastIndex := 0
	time := 0
	for lastIndex < len(nums)-1 {
		if lastIndex+nums[lastIndex] >= len(nums)-1 {
			lastIndex = len(nums) - 1
		} else {
			maxIndex := lastIndex + 1
			for i := lastIndex + 2; i < len(nums) && i <= lastIndex+nums[lastIndex]; i++ {
				if nums[i]+i > nums[maxIndex]+maxIndex {
					maxIndex = i
				}
			}
			lastIndex = maxIndex
		}
		time++
	}
	return time
}

func TestT45(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
	assert.Equal(t, 0, jump([]int{0}))
	assert.Equal(t, 0, jump([]int{0}))
	assert.Equal(t, 1, jump([]int{1, 2}))
	assert.Equal(t, 1, jump([]int{2, 1}))
	assert.Equal(t, 1, jump([]int{3, 2, 1}))
}
