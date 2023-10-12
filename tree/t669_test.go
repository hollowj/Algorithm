package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

func TestT669(t *testing.T) {
	assert.Equal(t, []int{1, null, 2}, trimBST(buildTree([]int{1, 0, 2}), 1, 2).FloorArray())
	assert.Equal(t, []int{3, 2, null, 1}, trimBST(buildTree([]int{3, 0, 4, null, 2, null, null, 1}), 1, 3).FloorArray())
	assert.Equal(t, []int{36, 34, 42, 33, 35, 41, 43, 32, null, null, null, 37, null, null, 44, null, null, null, 38, null, null, null, 39, null, 40}, trimBST(buildTree([]int{45, 30, 46, 10, 36, null, 49, 8, 24, 34, 42, 48, null, 4, 9, 14, 25, 31, 35, 41, 43, 47, null, 0, 6, null, null, 11, 20, null, 28, null, 33, null, null, 37, null, null, 44, null, null, null, 1, 5, 7, null, 12, 19, 21, 26, 29, 32, null, null, 38, null, null, null, 3, null, null, null, null, null, 13, 18, null, null, 22, null, 27, null, null, null, null, null, 39, 2, null, null, null, 15, null, null, 23, null, null, null, 40, null, null, null, 16, null, null, null, null, null, 17}), 32, 44).FloorArray())

}
