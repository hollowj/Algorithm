package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxArea(height []int) int {
	iLeft, iRight := 0, len(height)-1
	iMaxArea := 0
	for iLeft < iRight {
		var area int
		if height[iLeft] < height[iRight] {
			area = height[iLeft] * (iRight - iLeft)
			iLeft++

		} else {
			area = height[iRight] * (iRight - iLeft)
			iRight--
		}
		if area > iMaxArea {
			iMaxArea = area
		}

	}
	return iMaxArea
}

func TestT11(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 1, maxArea([]int{1, 1}))
}
