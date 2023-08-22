package single_stack

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

//func trap(height []int) int {
//	result := 0
//	lMax := make([]int, len(height))
//	lMax[0] = height[0]
//	for i := 1; i < len(height); i++ {
//		if height[i] > lMax[i-1] {
//			lMax[i] = height[i]
//		} else {
//			lMax[i] = lMax[i-1]
//
//		}
//	}
//	rMax := make([]int, len(height))
//	rMax[len(height)-1] = height[len(height)-1]
//	for i := len(height) - 2; i >= 0; i-- {
//		if height[i] > rMax[i+1] {
//			rMax[i] = height[i]
//		} else {
//			rMax[i] = rMax[i+1]
//
//		}
//	}
//	for i := 0; i < len(height); i++ {
//		h := min(lMax[i], rMax[i]) - height[i]
//		if h > 0 {
//			result += h
//		}
//	}
//	return result
//}

//func trap(height []int) int {
//	result := 0
//	stack := linkedliststack.New()
//	stack.Push(0)
//	for i := 1; i < len(height); i++ {
//		if height[i] < height[peekInt(stack)] {
//			stack.Push(i)
//		} else if height[i] == height[peekInt(stack)] {
//			stack.Pop()
//			stack.Push(i)
//		} else {
//			for !stack.Empty() && height[i] > height[peekInt(stack)] {
//				value, _ := stack.Pop()
//				mid := value.(int)
//				if !stack.Empty() {
//					left := peekInt(stack)
//					w := i - left - 1
//					h := min(height[left], height[i]) - height[mid]
//					iArea := w * h
//					if iArea > 0 {
//						result += iArea
//					}
//				}
//			}
//			stack.Push(i)
//		}
//
//	}
//	return result
//}

func trap(height []int) int {
	result := 0
	stack := linkedliststack.New()
	stack.Push(0)
	for i := 1; i < len(height); i++ {
		for !stack.Empty() && height[i] > height[peekInt(stack)] {
			value, _ := stack.Pop()
			mid := value.(int)
			if !stack.Empty() {
				left := peekInt(stack)
				w := i - left - 1
				h := min(height[left], height[i]) - height[mid]
				iArea := w * h
				if iArea > 0 {
					result += iArea
				}
			}
		}
		stack.Push(i)
	}

	return result
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func TestT42(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5}))
}
