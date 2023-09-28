package tree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func rightSideView(root *TreeNode) []int {
	results := make([]int, 0)
	nextFloor := []*TreeNode{}
	if root != nil {
		nextFloor = append(nextFloor, root)
	}
	for len(nextFloor) != 0 {
		tmp := make([]*TreeNode, 0)
		results = append(results, nextFloor[len(nextFloor)-1].Val)
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

	return results
}

func TestT199(t *testing.T) {
	assert.Equal(t, []int{1, 3, 4}, rightSideView(buildTree([]int{1, 2, 3, math.MinInt, 5, math.MinInt, 4})))
	assert.Equal(t, []int{1, 3}, rightSideView(buildTree([]int{1, math.MinInt, 3})))
	assert.Equal(t, []int{}, rightSideView(buildTree([]int{})))
	assert.Equal(t, []int{1, 3, 4}, rightSideView(buildTree([]int{1, 2, 3, 4})))
}
