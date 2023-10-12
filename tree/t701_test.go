package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
func TestT701(t *testing.T) {
	assert.Equal(t, []int{4, 2, 7, 1, 3, 5}, insertIntoBST(buildTree([]int{4, 2, 7, 1, 3}), 5).FloorArray())
	assert.Equal(t, []int{40, 20, 60, 10, 30, 50, 70, null, null, 25}, insertIntoBST(buildTree([]int{40, 20, 60, 10, 30, 50, 70}), 25).FloorArray())
	assert.Equal(t, []int{4, 2, 7, 1, 3, 5}, insertIntoBST(buildTree([]int{4, 2, 7, 1, 3, null, null, null, null, null, null}), 5).FloorArray())
	assert.Equal(t, []int{5}, insertIntoBST(buildTree([]int{}), 5).FloorArray())
	//assert.Equal(t, []int{0}, lowestCommonAncestor(buildTree([]int{0})))
	//assert.Equal(t, []int{2, 1}, lowestCommonAncestor(buildTree([]int{1, null, 2})))
	//assert.Equal(t, []int{0, 1}, lowestCommonAncestor(buildTree([]int{1, 0, 1, 0, 0, 1, 1, 0})))

}
