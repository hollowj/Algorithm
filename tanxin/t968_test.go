package tanxin

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
		return &TreeNode{}
	}
}

func buildTree(data []int) *TreeNode {
	root := &TreeNode{}
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

func minCameraCover(root *TreeNode) int {
	count := 0
	if dfs(root, &count) == 0 {
		count++
	}
	return count
}

// 0：该节点无覆盖
// 1：本节点有摄像头
// 2：本节点有覆盖
func dfs(root *TreeNode, count *int) int {
	if root == nil {
		return 2
	}
	left := dfs(root.Left, count)
	right := dfs(root.Right, count)
	if left == 2 && right == 2 {
		return 0
	}
	if left == 0 || right == 0 {
		*count += 1
		return 1
	}
	if left == 1 || right == 1 {
		return 2
	}
	return 0
}
func TestT968(t *testing.T) {
	assert.Equal(t, 1, minCameraCover(buildTree([]int{0, 0, -1, 0, 0})))

}
