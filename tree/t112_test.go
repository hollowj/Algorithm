package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func TestT112(t *testing.T) {
	assert.Equal(t, true, hasPathSum(buildTree([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1}), 22))
	assert.Equal(t, false, hasPathSum(buildTree([]int{1, 2, 3}), 5))

}
