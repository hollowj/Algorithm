package tree

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

func inorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	stack := linkedliststack.New()
	cur := root
	for cur != nil || !stack.Empty() {
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else {
			value, _ := stack.Pop()
			cur = value.(*TreeNode)
			if cur != nil {
				arr = append(arr, cur.Val)
				cur = cur.Right
			}
		}

	}
	return arr
}

//func inorderTraversal(root *TreeNode) []int {
//	arr := make([]int, 0)
//	inorderTraversal1(root, &arr)
//	return arr
//}
//func inorderTraversal1(root *TreeNode, arr *[]int) {
//	if root == nil {
//		return
//	}
//	inorderTraversal1(root.Left, arr)
//	*arr = append(*arr, root.Val)
//	inorderTraversal1(root.Right, arr)
//
//}

func TestT94(t *testing.T) {
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(buildTree([]int{1, -1, 2, 3})))
}
