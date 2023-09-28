package tree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func levelOrder(root *TreeNode) [][]int {
//	results := make([][]int, 0)
//	nextFloor := linkedlistqueue.New()
//	if root != nil {
//		nextFloor.Enqueue(root)
//	}
//	for !nextFloor.Empty() {
//		size := nextFloor.Size()
//		result := make([]int, 0)
//		for i := 0; i < size; i++ {
//			front, _ := nextFloor.Dequeue()
//			node := front.(*TreeNode)
//			result = append(result, node.Val)
//			if node.Left != nil {
//				nextFloor.Enqueue(node.Left)
//			}
//			if node.Right != nil {
//				nextFloor.Enqueue(node.Right)
//			}
//		}
//		results = append(results, result)
//	}
//	return results
//}

//func levelOrder(root *TreeNode) [][]int {
//	results := make([][]int, 0)
//	nextFloor := list.New()
//	if root != nil {
//		nextFloor.PushBack(root)
//	}
//	for nextFloor.Len() != 0 {
//		size := nextFloor.Len()
//		result := make([]int, 0)
//		for i := 0; i < size; i++ {
//			front := nextFloor.Front()
//			nextFloor.Remove(front)
//			node := front.Value.(*TreeNode)
//			result = append(result, node.Val)
//			if node.Left != nil {
//				nextFloor.PushBack(node.Left)
//			}
//			if node.Right != nil {
//				nextFloor.PushBack(node.Right)
//			}
//		}
//		results = append(results, result)
//	}
//	return results
//}

func levelOrder(root *TreeNode) [][]int {
	results := make([][]int, 0)
	nextFloor := []*TreeNode{}
	if root != nil {
		nextFloor = append(nextFloor, root)
	}
	for len(nextFloor) != 0 {
		result := make([]int, 0)
		tmp := make([]*TreeNode, 0)
		for _, node := range nextFloor {
			result = append(result, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		results = append(results, result)
		nextFloor = tmp
	}

	return results
}

//	func levelOrder(root *TreeNode) [][]int {
//		results := make([][]int, 0)
//		nextFloor := []*TreeNode{}
//		if root != nil {
//			nextFloor = append(nextFloor, root)
//		}
//		for len(nextFloor) != 0 {
//			size := len(nextFloor)
//			result := make([]int, 0)
//			for i := 0; i < size; i++ {
//				node := nextFloor[i]
//				result = append(result, node.Val)
//				if node.Left != nil {
//					nextFloor = append(nextFloor, node.Left)
//				}
//				if node.Right != nil {
//					nextFloor = append(nextFloor, node.Right)
//				}
//			}
//			results = append(results, result)
//			nextFloor = nextFloor[size:]
//		}
//		return results
//	}
func TestT102(t *testing.T) {
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, levelOrder(buildTree([]int{3, 9, 20, math.MinInt, math.MinInt, 15, 7})))
	assert.Equal(t, [][]int{{1}}, levelOrder(buildTree([]int{1})))
	assert.Equal(t, [][]int{}, levelOrder(buildTree([]int{})))
}
