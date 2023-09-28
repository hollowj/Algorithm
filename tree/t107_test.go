package tree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func levelOrderBottom(root *TreeNode) [][]int {
	results := make([][]int, 0)
	nextFloor := []*TreeNode{}
	if root != nil {
		nextFloor = append(nextFloor, root)
	}
	for len(nextFloor) != 0 {
		result := make([]int, 0)
		tmp := make([]*TreeNode, 0)
		for _, node := range nextFloor {
			result = append(result, node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		results = append(results, result)
		nextFloor = tmp
	}
	for i := 0; i < len(results)/2; i++ {
		results[i], results[len(results)-1-i] = results[len(results)-1-i], results[i]
	}
	return results
}

func TestT107(t *testing.T) {
	assert.Equal(t, [][]int{{15, 7}, {9, 20}, {3}}, levelOrderBottom(buildTree([]int{3, 9, 20, math.MinInt, math.MinInt, 15, 7})))
	assert.Equal(t, [][]int{{1}}, levelOrder(buildTree([]int{1})))
	assert.Equal(t, [][]int{}, levelOrder(buildTree([]int{})))
}
