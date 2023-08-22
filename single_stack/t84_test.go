package single_stack

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

//func largestRectangleArea(heights []int) int {
//	lLower := make([]int, len(heights))
//	lLower[0] = -1
//	for i := 1; i < len(heights); i++ {
//		t := i - 1
//		for t >= 0 && heights[t] >= heights[i] {
//			t = lLower[t]
//		}
//		lLower[i] = t
//	}
//	rLower := make([]int, len(heights))
//	rLower[len(heights)-1] = len(heights)
//	for i := len(heights) - 2; i >= 0; i-- {
//		t := i + 1
//		for t < len(heights) && heights[t] >= heights[i] {
//			t = rLower[t]
//		}
//		rLower[i] = t
//	}
//	iMaxArea := 0
//	for i := 0; i < len(heights); i++ {
//		left := lLower[i]
//		right := rLower[i]
//		iArea := (right - left - 1) * heights[i]
//		if iArea > iMaxArea {
//			iMaxArea = iArea
//		}
//	}
//	return iMaxArea
//}

//	func largestRectangleArea(heights []int) int {
//		iMaxArea := 0
//		stack := linkedliststack.New()
//		rightAreas := make([]int, len(heights))
//		for i := 0; i < len(heights)+1; i++ {
//			for !stack.Empty() && (i == len(heights) || heights[peekInt(stack)] > heights[i]) {
//				value, _ := stack.Pop()
//				left := value.(int)
//				area := (i - left) * heights[left]
//				rightAreas[left] = area
//
//			}
//			stack.Push(i)
//		}
//		stack.Clear()
//		leftAreas := make([]int, len(heights))
//		for i := len(heights) - 1; i >= -1; i-- {
//			for !stack.Empty() && (i == -1 || heights[peekInt(stack)] > heights[i]) {
//				value, _ := stack.Pop()
//				right := value.(int)
//				area := (right - i) * heights[right]
//				leftAreas[right] = area
//
//			}
//			stack.Push(i)
//		}
//
//		for i := 0; i < len(heights); i++ {
//			area := leftAreas[i] + rightAreas[i] - heights[i]
//			//area := rightAreas[i]
//			if area > iMaxArea {
//				iMaxArea = area
//			}
//		}
//		return iMaxArea
//	}
func largestRectangleArea(heights []int) int {
	iMaxArea := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack := linkedliststack.New()
	for i := 0; i < len(heights); i++ {
		for !stack.Empty() && heights[peekInt(stack)] > heights[i] {
			value, _ := stack.Pop()
			mid := value.(int)
			if !stack.Empty() {
				top := peekInt(stack)
				area := (i - top - 1) * heights[mid]
				if area > iMaxArea {
					iMaxArea = area
				}
			}
		}
		stack.Push(i)
	}

	return iMaxArea
}
func TestT84(t *testing.T) {
	assert.Equal(t, 10, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	assert.Equal(t, 4, largestRectangleArea([]int{2, 4}))
	assert.Equal(t, 2, largestRectangleArea([]int{1, 1}))
	assert.Equal(t, 3, largestRectangleArea([]int{2, 1, 2}))
}
