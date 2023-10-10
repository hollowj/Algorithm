package tree

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	type Node struct {
//		Val   int
//		Left  *Node
//		Right *Node
//		Next  *Node
//	}
//
//	func connect(root *Node) *Node {
//		nextFloor := []*Node{}
//		if root != nil {
//			nextFloor = append(nextFloor, root)
//		}
//		for len(nextFloor) != 0 {
//			tmp := make([]*Node, 0)
//			var pre *Node
//			for _, node := range nextFloor {
//				if pre != nil {
//					pre.Next = node
//				}
//				pre = node
//				if node.Left != nil {
//					tmp = append(tmp, node.Left)
//				}
//				if node.Right != nil {
//					tmp = append(tmp, node.Right)
//				}
//			}
//			nextFloor = tmp
//
//		}
//
//		return root
//	}

func binaryTreePaths(root *TreeNode) []string {
	results := make([]string, 0)
	dfs(root, "", &results)
	return results
}
func dfs(root *TreeNode, prefix string, results *[]string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*results = append(*results, prefix+strconv.Itoa(root.Val))
	}
	temp := prefix + strconv.Itoa(root.Val) + "->"
	dfs(root.Left, temp, results)
	dfs(root.Right, temp, results)

}
func TestT257(t *testing.T) {
	assert.Equal(t, []string{"1->2->5", "1->3"}, binaryTreePaths(buildTree([]int{1, 2, 3, null, 5})))
	assert.Equal(t, []string{"1"}, binaryTreePaths(buildTree([]int{1})))
}
