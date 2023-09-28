package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func averageOfLevels(root *TreeNode) []float64 {
	results := make([]float64, 0)
	nextFloor := []*TreeNode{}
	if root != nil {
		nextFloor = append(nextFloor, root)
	}
	for len(nextFloor) != 0 {
		tmp := make([]*TreeNode, 0)
		iSum := 0
		for _, node := range nextFloor {
			iSum += node.Val
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		results = append(results, float64(iSum)/float64(len(nextFloor)))
		nextFloor = tmp

	}

	return results
}

func TestT637(t *testing.T) {
	assert.Equal(t, []float64{3.00000, 14.50000, 11.00000}, averageOfLevels(buildTree([]int{3, 9, 20, null, null, 15, 7})))
	assert.Equal(t, []float64{3.00000, 14.50000, 11.00000}, averageOfLevels(buildTree([]int{3, 9, 20, 15, 7})))
}
