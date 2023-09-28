package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func minDepth(root *TreeNode) int {
	nextFloor := []*TreeNode{}
	if root != nil {
		nextFloor = append(nextFloor, root)
	}
	iDepth := 0
	for len(nextFloor) != 0 {
		iDepth++
		tmp := make([]*TreeNode, 0)
		for _, node := range nextFloor {
			if node.Left == nil && node.Right == nil {
				return iDepth
			}
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

func TestT111(t *testing.T) {
	assert.Equal(t, 2, minDepth(buildTree([]int{3, 9, 20, null, null, 15, 7})))
	assert.Equal(t, 5, minDepth(buildTree([]int{2, null, 3, null, 4, null, 5, null, 6})))
	assert.Equal(t, 3, minDepth(buildTree([]int{1, 2, 3, 4, null, null, 5})))
}
