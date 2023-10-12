package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	}
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	return nil
}

func TestT236(t *testing.T) {
	assert.Equal(t, []int{2}, lowestCommonAncestor(buildTree([]int{3, 5, 1, 6, 2, 0, 8, null, null, 7, 4}), buildTree([]int{5}), buildTree([]int{1})))
	//assert.Equal(t, []int{0}, lowestCommonAncestor(buildTree([]int{0})))
	//assert.Equal(t, []int{2, 1}, lowestCommonAncestor(buildTree([]int{1, null, 2})))
	//assert.Equal(t, []int{0, 1}, lowestCommonAncestor(buildTree([]int{1, 0, 1, 0, 0, 1, 1, 0})))

}
