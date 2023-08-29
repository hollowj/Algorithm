package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		pos, ok := m[target-num]
		if ok {
			return []int{i, pos}
		}
		m[num] = i
	}
	return []int{}
}

func TestT1(t *testing.T) {
	assert.Equal(t, []int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.Equal(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.Equal(t, []int{0, 1}, twoSum([]int{3, 3}, 6))
}
