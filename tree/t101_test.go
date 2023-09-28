package tree

import (
	"fmt"
	"testing"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/stretchr/testify/assert"
)

// func isSymmetric(root *TreeNode) bool {
//
//		nextFloor := []*TreeNode{}
//		if root != nil {
//			nextFloor = append(nextFloor, root)
//		}
//		for len(nextFloor) != 0 {
//			tmp := make([]*TreeNode, 0)
//			iNilCount := 0
//			for _, node := range nextFloor {
//				if node != nil {
//					tmp = append(tmp, node.Left)
//					tmp = append(tmp, node.Right)
//				} else {
//					tmp = append(tmp, nil)
//					tmp = append(tmp, nil)
//					iNilCount++
//				}
//
//			}
//			if iNilCount == len(nextFloor) {
//				break
//			}
//			for i := 0; i < len(tmp)/2; i++ {
//				before := tmp[i]
//				after := tmp[len(tmp)-1-i]
//				if before == nil || after == nil {
//					if before != after {
//						return false
//					} else {
//						continue
//					}
//				}
//				if before.Val != after.Val {
//					return false
//				}
//			}
//
//			nextFloor = tmp
//
//		}
//
//		return true
//	}
//
//	func isSymmetric(root *TreeNode) bool {
//		if root == nil {
//			return true
//		}
//		return compare(root.Left, root.Right)
//	}
//
//	func compare(left *TreeNode, right *TreeNode) bool {
//		if left == nil || right == nil {
//			if left == nil {
//				if left != right {
//					return false
//				} else {
//					return true
//				}
//			} else {
//				return false
//			}
//		}
//		if left.Val != right.Val {
//			return false
//		}
//		isInnerSame := compare(left.Right, right.Left)
//		isOutterSame := compare(left.Left, right.Right)
//		return isInnerSame && isOutterSame
//	}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := linkedlistqueue.New()
	queue.Enqueue(root.Left)
	queue.Enqueue(root.Right)
	for !queue.Empty() {
		left, _ := queue.Dequeue()
		right, _ := queue.Dequeue()
		treeNode1 := left.(*TreeNode)
		treeNode2 := right.(*TreeNode)
		if treeNode1 == nil || treeNode2 == nil {
			if treeNode1 == nil {
				if treeNode1 != treeNode2 {
					return false
				} else {
					continue
				}
			} else {
				return false
			}
		}

		if treeNode1.Val != treeNode2.Val {
			return false
		}
		queue.Enqueue(treeNode1.Left)
		queue.Enqueue(treeNode2.Right)
		queue.Enqueue(treeNode1.Right)
		queue.Enqueue(treeNode2.Left)
	}
	return true
}
func TestT101(t *testing.T) {
	assert.Equal(t, true, isSymmetric(buildTree([]int{1, 2, 2, 3, 4, 4, 3})))
	assert.Equal(t, false, isSymmetric(buildTree([]int{1, 2, 2, null, 3, null, 3})))
	assert.Equal(t, false, isSymmetric(buildTree([]int{9, -42, -42, null, 76, 76, null, null, 13, null, 13})))
}

func TestNil(t *testing.T) {
	node := NewTreeNode(1)
	arr := make([]interface{}, 0)
	arr = append(arr, node)
	arr = append(arr, node.Right)
	for _, val := range arr {
		treeNode := val.(*TreeNode)
		if treeNode == nil {
			fmt.Println("exist nil")
		}
	}
}
