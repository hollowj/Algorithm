package single_stack

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

func nextGreaterElements(nums []int) []int {
	stack := linkedliststack.New()
	size := len(nums)
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = -1
	}
	for i := 0; i < 2*size; i++ {
		for !stack.Empty() && nums[peekInt(stack)%size] < nums[i%size] {
			value, _ := stack.Pop()
			v := value.(int)
			if result[v%size] == -1 {
				result[v%size] = nums[i%size]
			}
		}
		stack.Push(i)
	}
	return result
}

func TestT503(t *testing.T) {
	assert.Equal(t, []int{2, -1, 2}, nextGreaterElements([]int{1, 2, 1}))
	assert.Equal(t, []int{2, 3, 4, -1, 4}, nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
