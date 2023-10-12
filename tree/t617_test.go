package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
//		if root1 == nil && root2 == nil {
//			return nil
//		}
//		var cur = root1
//		var left1 *TreeNode
//		var left2 *TreeNode
//		var right1 *TreeNode
//		var right2 *TreeNode
//		if root1 != nil {
//			left1 = root1.Left
//			right1 = root1.Right
//			if root2 != nil {
//				root1.Val += root2.Val
//				left2 = root2.Left
//				right2 = root2.Right
//			}
//		} else {
//			cur = root2
//			left2 = root2.Left
//			right2 = root2.Right
//		}
//		cur.Left = mergeTrees(left1, left2)
//		cur.Right = mergeTrees(right1, right2)
//		return cur
//	}
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func TestT617(t *testing.T) {
	assert.Equal(t, []int{3, 4, 5, 5, 4, null, 7}, mergeTrees(buildTree([]int{1, 3, 2, 5}), buildTree([]int{2, 1, 3, null, 4, null, 7})).FloorArray())
	assert.Equal(t, []int{2, 2}, mergeTrees(buildTree([]int{1}), buildTree([]int{1, 2})).FloorArray())

}
