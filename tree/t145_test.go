package tree

import (
	"container/list"
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

func postorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0)
	list.New()
	stack := linkedliststack.New()
	if root != nil {
		stack.Push(root)
	}
	for !stack.Empty() {
		value, _ := stack.Peek()
		if value != nil {
			stack.Pop()
			cur := value.(*TreeNode)
			stack.Push(cur)
			stack.Push(nil)
			if cur.Right != nil {
				stack.Push(cur.Right)
			}
			if cur.Left != nil {
				stack.Push(cur.Left)
			}
		} else {
			stack.Pop()
			value, _ = stack.Pop()
			cur := value.(*TreeNode)
			arr = append(arr, cur.Val)

		}

	}
	return arr
}

//
//func postorderTraversal(root *TreeNode) []int {
//	arr := make([]int, 0)
//	stack := linkedliststack.New()
//	stack.Push(root)
//	for !stack.Empty() {
//		value, _ := stack.Pop()
//		treeNode := value.(*TreeNode)
//		arr = append(arr, treeNode.Val)
//		if treeNode.Left != nil {
//			stack.Push(treeNode.Left)
//		}
//		if treeNode.Right != nil {
//			stack.Push(treeNode.Right)
//		}
//	}
//	reverse(arr)
//	return arr
//}
//
//func reverse(a []int) {
//	l, r := 0, len(a)-1
//	for l < r {
//		a[l], a[r] = a[r], a[l]
//		l, r = l+1, r-1
//	}
//}

//func postorderTraversal(root *TreeNode) []int {
//	arr := make([]int, 0)
//	postorderTraversal1(root, &arr)
//	return arr
//}
//func postorderTraversal1(root *TreeNode, arr *[]int) {
//	if root == nil {
//		return
//	}
//	postorderTraversal1(root.Left, arr)
//	postorderTraversal1(root.Right, arr)
//	*arr = append(*arr, root.Val)
//
//}

func TestT145(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, postorderTraversal(buildTree([]int{1, -1, 2, 3})))
}
