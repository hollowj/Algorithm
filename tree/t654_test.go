package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	iMaxIndex := 0
	for i, num := range nums {
		if num > nums[iMaxIndex] {
			iMaxIndex = i
		}
	}
	root := &TreeNode{Val: nums[iMaxIndex]}
	root.Left = constructMaximumBinaryTree(nums[:iMaxIndex])
	root.Right = constructMaximumBinaryTree(nums[iMaxIndex+1:])
	return root
}

func TestT654(t *testing.T) {
	assert.Equal(t, []int{6, 3, 5, null, 2, 0, null, null, 1}, constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5}).FloorArray())
	assert.Equal(t, []int{3, null, 2, null, 1}, constructMaximumBinaryTree([]int{3, 2, 1}).FloorArray())

}
