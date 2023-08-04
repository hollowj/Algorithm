package tanxin

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	maxDistance := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxDistance < i {
			return false
		}
		maxDistance = int(math.Max(float64(maxDistance), float64(i+nums[i])))
	}
	return true
}
func TestT55(t *testing.T) {
	assert.Equal(t, true, canJump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, false, canJump([]int{3, 2, 1, 0, 4}))
}
