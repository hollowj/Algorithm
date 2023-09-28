package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	return IsSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
func IsSameTree(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	isInnerSame := IsSameTree(left.Left, right.Left)
	isOutterSame := IsSameTree(left.Right, right.Right)
	return isInnerSame && isOutterSame
}

func TestT572(t *testing.T) {
	assert.Equal(t, true, isSubtree(buildTree([]int{3, 4, 5, 1, 2}), buildTree([]int{4, 1, 2})))
	assert.Equal(t, false, isSubtree(buildTree([]int{3, 4, 5, 1, 2, null, null, null, null, 0}), buildTree([]int{4, 1, 2})))
}
