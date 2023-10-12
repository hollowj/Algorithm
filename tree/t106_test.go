package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	mid := postorder[len(postorder)-1]
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
	postorder = postorder[:len(postorder)-1]
	leftPostorder := postorder[:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder):]
	root.Left = buildTree106(leftInorder, leftPostorder)
	root.Right = buildTree106(rightInorder, rightPostorder)
	return root
}

//	func buildTree106(inorder []int, postorder []int) *TreeNode {
//		return buildTheTree(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
//	}
//
//	func buildTheTree(inorder []int, postorder []int, inorderStart, inorderEnd, postorderStart, postorderEnd int) *TreeNode {
//		if inorderStart > inorderEnd {
//			return nil
//		}
//		mid := postorder[postorderEnd]
//		root := &TreeNode{Val: mid}
//		sepIndex := 0
//		for i, val := range inorder {
//			if val == mid {
//				sepIndex = i
//				break
//			}
//		}
//		size := sepIndex - inorderStart
//		//leftInorder := inorder[:sepIndex]
//		//rightInorder := inorder[sepIndex+1:]
//		//postorder = postorder[:len(postorder)-1]
//		//leftPostorder := postorder[:len(leftInorder)]
//		//rightPostorder := postorder[len(leftInorder):]
//		root.Left = buildTheTree(inorder, postorder, inorderStart, sepIndex-1, postorderStart, postorderStart+size-1)
//		root.Right = buildTheTree(inorder, postorder, sepIndex+1, inorderEnd, postorderStart+size, postorderEnd-1)
//		return root
//	}
func TestT106(t *testing.T) {
	assert.Equal(t, []int{3, 9, 20, null, null, 15, 7}, buildTree106([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}).FloorArray())
	assert.Equal(t, []int{-1}, buildTree106([]int{-1}, []int{-1}).FloorArray())

}
