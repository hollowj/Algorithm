package single_stack

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := linkedliststack.New()
	m := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		for !stack.Empty() && nums2[peekInt(stack)] < nums2[i] {
			value, _ := stack.Pop()
			index := value.(int)
			m[nums2[index]] = nums2[i]
		}
		stack.Push(i)
	}
	result1 := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		v, ok := m[nums1[i]]
		if ok {
			result1[i] = v
		} else {
			result1[i] = -1

		}
	}
	return result1
}
func TestT496(t *testing.T) {
	assert.Equal(t, []int{-1, 3, -1}, nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	assert.Equal(t, []int{3, -1}, nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
	assert.Equal(t, []int{-1, -1, -1, -1, -1}, nextGreaterElement([]int{1, 3, 5, 2, 4}, []int{5, 4, 3, 2, 1}))
}
