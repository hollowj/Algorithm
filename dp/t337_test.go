package dp

import (
	"testing"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTreeNode(n int) *TreeNode {
	if n < 0 {
		return nil
	} else {
		return &TreeNode{Val: n}
	}
}

func buildTree(data []int) *TreeNode {
	root := newTreeNode(data[0])
	queue := linkedlistqueue.New()
	queue.Enqueue(root)
	index := 0
	for index < len(data) {
		tmp := linkedlistqueue.New()
		for queue.Size() > 0 {
			value, _ := queue.Dequeue()
			node := value.(*TreeNode)
			index++
			if index < len(data) {
				node.Left = newTreeNode(data[index])
				if node.Left != nil {
					tmp.Enqueue(node.Left)
				}
			} else {
				break
			}
			index++
			if index < len(data) {
				node.Right = newTreeNode(data[index])
				if node.Right != nil {
					tmp.Enqueue(node.Right)
				}
			}
		}
		queue = tmp
	}
	return root
}
func rob337(root *TreeNode) int {
	t := rob_r(root)
	return max(t[0], t[1])
}
func rob_r(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	lGain := rob_r(root.Left)
	rGain := rob_r(root.Right)
	return []int{max(lGain[0], lGain[1]) + max(rGain[0], rGain[1]), lGain[0] + rGain[0] + root.Val}
}
func TestT337(t *testing.T) {
	assert.Equal(t, 7, rob337(buildTree([]int{3, 2, 3, -1, 3, -1, 1})))
	assert.Equal(t, 9, rob337(buildTree([]int{3, 4, 5, 1, 3, -1, 1})))
}
