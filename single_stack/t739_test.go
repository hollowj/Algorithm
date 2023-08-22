package single_stack

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

func dailyTemperatures(temperatures []int) []int {
	stack := linkedliststack.New()
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for !stack.Empty() && temperatures[peekInt(stack)] < temperatures[i] {
			value, _ := stack.Pop()
			result[value.(int)] = i - value.(int)
		}
		stack.Push(i)
	}
	return result
}

func peekInt(stack *linkedliststack.Stack) int {
	i, _ := stack.Peek()
	return i.(int)
}
func TestT739(t *testing.T) {
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	assert.Equal(t, []int{1, 1, 1, 0}, dailyTemperatures([]int{30, 40, 50, 60}))
	assert.Equal(t, []int{1, 1, 0}, dailyTemperatures([]int{30, 60, 90}))
}
