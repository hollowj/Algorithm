package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func TestT108(t *testing.T) {
	assert.Equal(t, []int{0, -3, 9, -10, null, 5}, sortedArrayToBST([]int{-10, -3, 0, 5, 9}).FloorArray())
	assert.Equal(t, []int{3, 1}, sortedArrayToBST([]int{1, 3}).FloorArray())

}
