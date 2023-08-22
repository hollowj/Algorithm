package single_stack

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

func peekInt(stack *linkedliststack.Stack) int {
	i, _ := stack.Peek()
	return i.(int)
}
func trap(height []int) int {
	stack := linkedliststack.New()
	stack.Push(0)
	result := 0
	for i := 1; i < len(height); i++ {
		for !stack.Empty() && height[peekInt(stack)] < height[i] {
			value, _ := stack.Pop()
			mid := value.(int)
			if !stack.Empty() {
				top := peekInt(stack)
				result += (i - top - 1) * (min(height[i], height[top]) - height[mid])
			}
		}
		stack.Push(i)
	}
	return result
}

// 暴力遍历
//
//	func trap(height []int) int {
//		result := 0
//		for i := 1; i < len(height)-1; i++ {
//			lHeightestIndex := i - 1
//			rHeightestIndex := i + 1
//			for j := i - 1; j >= 0; j-- {
//				if height[j] > height[lHeightestIndex] {
//					lHeightestIndex = j
//				}
//			}
//			for j := i + 1; j < len(height); j++ {
//				if height[j] > height[rHeightestIndex] {
//					rHeightestIndex = j
//				}
//			}
//			h := min(height[lHeightestIndex], height[rHeightestIndex]) - height[i]
//			if h > 0 {
//				result += h
//			}
//		}
//		return result
//	}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 双指针
//
//	func trap(height []int) int {
//		result := 0
//		leftMax := make([]int, len(height))
//		rightMax := make([]int, len(height))
//		leftMax[0] = height[0]
//		for i := 1; i < len(height); i++ {
//			leftMax[i] = max(leftMax[i-1], height[i])
//		}
//		rightMax[len(height)-1] = height[len(height)-1]
//		for i := len(height) - 2; i > 0; i-- {
//			rightMax[i] = max(rightMax[i+1], height[i])
//		}
//		for i := 0; i < len(height); i++ {
//			if i == 0 || i == len(height)-1 {
//				continue
//			}
//
//			h := min(leftMax[i], rightMax[i]) - height[i]
//			if h > 0 {
//				result += h
//			}
//
//		}
//		return result
//	}
func TestT42(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5}))
}
