package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	ileftSum := sumOfLeftLeaves(root.Left)
	iRightSum := sumOfLeftLeaves(root.Right)
	return ileftSum + iRightSum
}

func TestT404(t *testing.T) {
	assert.Equal(t, 24, sumOfLeftLeaves(buildTree([]int{3, 9, 20, null, null, 15, 7})))
	assert.Equal(t, 0, sumOfLeftLeaves(buildTree([]int{1})))
}
