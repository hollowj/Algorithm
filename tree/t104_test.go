package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxDepth(root *TreeNode) int {
	nextFloor := []*TreeNode{}
	if root != nil {
		nextFloor = append(nextFloor, root)
	}
	iDepth := 0
	for len(nextFloor) != 0 {
		iDepth++
		tmp := make([]*TreeNode, 0)
		for _, node := range nextFloor {
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		nextFloor = tmp

	}

	return iDepth
}

func TestT104(t *testing.T) {
	assert.Equal(t, 3, maxDepth(buildTree([]int{3, 9, 20, null, null, 15, 7})))
	assert.Equal(t, 2, maxDepth(buildTree([]int{1, null, 2})))
}
