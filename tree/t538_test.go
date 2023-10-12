package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func convertBST(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode)
	var pre *TreeNode
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		if pre != nil {
			root.Val += pre.Val
		}
		pre = root
		dfs(root.Left)
	}
	dfs(root)
	return root
}

func TestT538(t *testing.T) {
	assert.Equal(t, []int{30, 36, 21, 36, 35, 26, 15, null, null, null, 33, null, null, null, 8}, convertBST(buildTree([]int{4, 1, 6, 0, 2, 5, 7, null, null, null, 3, null, null, null, 8})).FloorArray())
	assert.Equal(t, []int{1, null, 1}, convertBST(buildTree([]int{0, null, 1})).FloorArray())
	assert.Equal(t, []int{3, 3, 2}, convertBST(buildTree([]int{1, 0, 2})).FloorArray())
	assert.Equal(t, []int{7, 9, 4, 10}, convertBST(buildTree([]int{3, 2, 4, 1})).FloorArray())

}
