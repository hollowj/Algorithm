package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		if p == nil {
			if p != q {
				return false
			} else {
				return true
			}
		} else {
			return false
		}
	}
	if p.Val != q.Val {
		return false
	}
	isInnerSame := isSameTree(p.Left, q.Left)
	isOuterSame := isSameTree(p.Right, q.Right)
	return isInnerSame && isOuterSame
}

func TestT100(t *testing.T) {
	assert.Equal(t, true, isSameTree(buildTree([]int{1, 2, 3}), buildTree([]int{1, 2, 3})))
	assert.Equal(t, false, isSameTree(buildTree([]int{1, 2}), buildTree([]int{1, null, 2})))
	assert.Equal(t, false, isSameTree(buildTree([]int{1, 2, 1}), buildTree([]int{1, 1, 2})))
}
