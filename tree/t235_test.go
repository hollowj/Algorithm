package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	var dfs func(root, p, q *TreeNode) *TreeNode
	dfs = func(root, p, q *TreeNode) *TreeNode {
		if root == nil {
			return root
		}
		if root.Val > p.Val && root.Val > q.Val {
			l := dfs(root.Left, p, q)
			if l != nil {
				return l
			}
		}
		if root.Val < p.Val && root.Val < q.Val {
			r := dfs(root.Right, p, q)
			if r != nil {
				return r
			}
		}
		return root
	}
	return dfs(root, p, q)
}
func TestT235(t *testing.T) {
	assert.Equal(t, []int{2}, lowestCommonAncestor235(buildTree([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}), buildTree([]int{2}), buildTree([]int{8})))
	assert.Equal(t, []int{2}, lowestCommonAncestor235(buildTree([]int{6, 2, 8, 0, 4, 7, 9, null, null, 3, 5}), buildTree([]int{2}), buildTree([]int{4})))
	//assert.Equal(t, []int{0}, lowestCommonAncestor(buildTree([]int{0})))
	//assert.Equal(t, []int{2, 1}, lowestCommonAncestor(buildTree([]int{1, null, 2})))
	//assert.Equal(t, []int{0, 1}, lowestCommonAncestor(buildTree([]int{1, 0, 1, 0, 0, 1, 1, 0})))

}
