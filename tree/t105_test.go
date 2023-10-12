package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildTree105(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	mid := preorder[0]
	root := &TreeNode{Val: mid}
	sepIndex := 0
	for i, val := range inorder {
		if val == mid {
			sepIndex = i
			break
		}
	}
	leftInorder := inorder[:sepIndex]
	rightInorder := inorder[sepIndex+1:]
	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1:]
	root.Left = buildTree105(leftPreorder, leftInorder)
	root.Right = buildTree105(rightPreorder, rightInorder)
	return root
}

func TestT105(t *testing.T) {
	assert.Equal(t, []int{3, 9, 20, null, null, 15, 7}, buildTree105([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}).FloorArray())
	assert.Equal(t, []int{-1}, buildTree105([]int{-1}, []int{-1}).FloorArray())

}
