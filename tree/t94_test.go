package tree

import (
	"fmt"
	"math"
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

func inorderTraversal(root *TreeNode) []int {
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
			stack.Push(cur)
			stack.Push(nil)
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

//func inorderTraversal(root *TreeNode) []int {
//	arr := make([]int, 0)
//	stack := linkedliststack.New()
//	cur := root
//	for cur != nil || !stack.Empty() {
//		if cur != nil {
//			stack.Push(cur)
//			cur = cur.Left
//		} else {
//			value, _ := stack.Pop()
//			cur = value.(*TreeNode)
//			if cur != nil {
//				arr = append(arr, cur.Val)
//				cur = cur.Right
//			}
//		}
//
//	}
//	return arr
//}

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
	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(buildTree([]int{1, math.MinInt, 2, 3})))
	assert.Equal(t, []int{}, inorderTraversal(buildTree([]int{})))
	assert.Equal(t, []int{1}, inorderTraversal(buildTree([]int{1})))
	fmt.Println(preorderTraversal(buildTree([]int{3, 9, 20, 6, 12, 15, 7})))
	fmt.Println(inorderTraversal(buildTree([]int{3, 9, 20, 6, 12, 15, 7})))
	fmt.Println(postorderTraversal(buildTree([]int{3, 9, 20, 6, 12, 15, 7})))
}

//
//func TestHttps(t *testing.T) {
//	response, err := http.Get("https://devops.partying.sg/project/game-center-china&game")
//	if err != nil {
//		panic(err)
//	}
//	bytes, err := io.ReadAll(response.Body)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(len(bytes))
//}
