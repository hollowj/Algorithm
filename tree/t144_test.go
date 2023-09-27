package tree

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

func preorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	stack := linkedliststack.New()
	if root != nil {
		stack.Push(root)
	}
	for !stack.Empty() {
		value, _ := stack.Peek()
		if value != nil {
			stack.Pop()
			cur := value.(*TreeNode)
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
			stack.Push(cur)
			stack.Push(nil)
		} else {
			stack.Pop()
			value, _ = stack.Pop()
			cur := value.(*TreeNode)
			arr = append(arr, cur.Val)

		}

	}
	return arr
}

//func preorderTraversal(root *TreeNode) []int {
//	arr := make([]int, 0)
//	stack := linkedliststack.New()
//	stack.Push(root)
//	for !stack.Empty() {
//		value, _ := stack.Pop()
//		treeNode := value.(*TreeNode)
//		arr = append(arr, treeNode.Val)
//		if treeNode.Right != nil {
//			stack.Push(treeNode.Right)
//		}
//		if treeNode.Left != nil {
//			stack.Push(treeNode.Left)
//		}
//	}
//	return arr
//}

//func preorderTraversal(root *TreeNode) []int {
//	arr := make([]int, 0)
//	preorderTraversal1(root, &arr)
//	return arr
//}
//func preorderTraversal1(root *TreeNode, arr *[]int) {
//	if root == nil {
//		return
//	}
//	*arr = append(*arr, root.Val)
//	preorderTraversal1(root.Left, arr)
//	preorderTraversal1(root.Right, arr)
//
//}

func TestT144(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, preorderTraversal(buildTree([]int{1, -1, 2, 3})))
}
