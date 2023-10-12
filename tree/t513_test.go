package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var iMaxDep int
var iVal int

func findBottomLeftValue(root *TreeNode) int {
	iMaxDep = 0
	iVal = 0
	findBottomLeft(root, 0)
	return iVal
}
func findBottomLeft(root *TreeNode, dep int) {
	if root == nil {
		return
	}
	dep = dep + 1
	if root.Left == nil && root.Right == nil {
		if dep > iMaxDep {
			iMaxDep = dep
			iVal = root.Val
		}
	}
	findBottomLeft(root.Left, dep)
	findBottomLeft(root.Right, dep)
	return
}

func TestT513(t *testing.T) {
	assert.Equal(t, 1, findBottomLeftValue(buildTree([]int{2, 1, 3})))
	assert.Equal(t, 7, findBottomLeftValue(buildTree([]int{1, 2, 3, 4, null, 5, 6, null, null, 7})))
	assert.Equal(t, 1, findBottomLeftValue(buildTree([]int{1})))
	assert.Equal(t, 6, findBottomLeftValue(buildTree([]int{10, 5, 15, null, null, 6, 20})))
	assert.Equal(t, 1, findBottomLeftValue(buildTree([]int{1, null, 1})))
}
