package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	tmp := searchBST(root.Left, val)
	if tmp != nil {
		return tmp
	}
	tmp = searchBST(root.Right, val)
	return tmp
}

func TestT700(t *testing.T) {
	assert.Equal(t, []int{2, 1, 3}, searchBST(buildTree([]int{4, 2, 7, 1, 3}), 2).FloorArray())
	assert.Equal(t, []int{}, searchBST(buildTree([]int{4, 2, 7, 1, 3}), 5).FloorArray())

}
