package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var pre *TreeNode

func isValidBST(root *TreeNode) bool {
	pre = nil
	return isValidBST1(root)
}
func isValidBST1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isValidBST1(root.Left) {
		return false
	}
	if pre != nil && pre.Val >= root.Val {
		return false
	}
	pre = root
	if !isValidBST1(root.Right) {
		return false
	}
	return true
}

func TestT98(t *testing.T) {
	assert.Equal(t, true, isValidBST(buildTree([]int{2, 1, 3})))
	assert.Equal(t, false, isValidBST(buildTree([]int{5, 1, 4, null, null, 3, 6})))
	assert.Equal(t, false, isValidBST(buildTree([]int{5, 4, 6, null, null, 3, 7})))
	assert.Equal(t, false, isValidBST(buildTree([]int{120, 70, 140, 50, 100, 130, 160, 20, 55, 75, 110, 119, 135, 150, 200})))

}
