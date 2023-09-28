package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func invertTree(root *TreeNode) *TreeNode {
//		nextFloor := []*TreeNode{}
//		if root != nil {
//			nextFloor = append(nextFloor, root)
//		}
//		for len(nextFloor) != 0 {
//			tmp := make([]*TreeNode, 0)
//			for _, node := range nextFloor {
//				if node.Left != nil {
//					tmp = append(tmp, node.Left)
//				}
//				if node.Right != nil {
//					tmp = append(tmp, node.Right)
//				}
//				node.Left, node.Right = node.Right, node.Left
//
//			}
//			nextFloor = tmp
//
//		}
//
//		return root
//	}
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
func TestT226(t *testing.T) {
	assert.Equal(t, []int{4, 7, 2, 9, 6, 3, 1}, invertTree(buildTree([]int{4, 2, 7, 1, 3, 6, 9})))
	assert.Equal(t, []float64{2, 3, 1}, invertTree(buildTree([]int{2, 1, 3})))
	assert.Equal(t, []float64{}, invertTree(buildTree([]int{})))
}
