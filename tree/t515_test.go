package tree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func largestValues(root *TreeNode) []int {
	results := make([]int, 0)
	nextFloor := []*TreeNode{}
	if root != nil {
		nextFloor = append(nextFloor, root)
	}
	for len(nextFloor) != 0 {
		tmp := make([]*TreeNode, 0)
		iMax := math.MinInt
		for _, node := range nextFloor {
			if node.Val > iMax {
				iMax = node.Val
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		results = append(results, iMax)
		nextFloor = tmp
	}

	return results
}

func TestT515(t *testing.T) {
	assert.Equal(t, []int{1, 3, 9}, largestValues(buildTree([]int{1, 3, 2, 5, 3, null, 9})))
	assert.Equal(t, []int{1, 3}, largestValues(buildTree([]int{1, 2, 3})))
	assert.Equal(t, []int{}, largestValues(buildTree([]int{})))
}
