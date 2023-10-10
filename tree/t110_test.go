package tree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isBalanced(root *TreeNode) bool {
	return getDepth(root) != -1
}
func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	iLeftDepth := getDepth(root.Left)
	if iLeftDepth == -1 {
		return -1
	}
	iRightDepth := getDepth(root.Right)
	if iRightDepth == -1 {
		return -1
	}
	if math.Abs(float64(iLeftDepth-iRightDepth)) > 1 {
		return -1
	}
	return max(iLeftDepth, iRightDepth) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestT110(t *testing.T) {
	assert.Equal(t, true, isBalanced(buildTree([]int{3, 9, 20, null, null, 15, 7})))
	assert.Equal(t, false, isBalanced(buildTree([]int{1, 2, 2, 3, 3, null, null, 4, 4})))
	assert.Equal(t, true, isBalanced(buildTree([]int{})))
	assert.Equal(t, false, isBalanced(buildTree([]int{1, 2, 2, 3, null, null, 3, 4, null, null, 4})))
}
