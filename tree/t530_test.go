package tree

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pre530 *TreeNode
var minValue int

func getMinimumDifference(root *TreeNode) int {
	pre530 = nil
	minValue = math.MaxInt
	dfs530(root)
	return minValue
}
func dfs530(root *TreeNode) {
	if root == nil {
		return
	}
	dfs530(root.Left)
	if pre530 != nil {
		iDiff := root.Val - pre530.Val
		if iDiff < minValue {
			minValue = iDiff
		}
	}
	pre530 = root
	dfs530(root.Right)

}

func TestT530(t *testing.T) {
	assert.Equal(t, 1, getMinimumDifference(buildTree([]int{4, 2, 6, 1, 3})))
	assert.Equal(t, 1, getMinimumDifference(buildTree([]int{1, 0, 48, null, null, 12, 49})))
	assert.Equal(t, 9, getMinimumDifference(buildTree([]int{236, 104, 701, null, 227, null, 911})))

}
