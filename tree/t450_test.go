package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key == root.Val {
		if root.Right != nil {
			l := root.Left
			t := root.Right
			for t != nil && t.Left != nil {
				t = t.Left
			}
			t.Left = l
			return root.Right
		} else {
			return root.Left
		}
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	return root
}
func TestT450(t *testing.T) {
	assert.Equal(t, []int{5, 4, 6, 2, null, null, 7}, deleteNode(buildTree([]int{5, 3, 6, 2, 4, null, 7}), 3).FloorArray())
	assert.Equal(t, []int{5, 3, 6, 2, 4, null, 7}, deleteNode(buildTree([]int{5, 3, 6, 2, 4, null, 7}), 0).FloorArray())

}
