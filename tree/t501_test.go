package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMode(root *TreeNode) []int {
	var dfs func(root *TreeNode)
	mostNum := []int{root.Val}
	var mostCount, lastCount int
	var preNode *TreeNode
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if preNode != nil {
			if root.Val == preNode.Val {
				lastCount++
			} else {
				lastCount = 1
			}
		} else {
			lastCount = 1
		}
		if lastCount > mostCount {
			mostCount = lastCount
			mostNum = []int{root.Val}
		} else if lastCount == mostCount {
			mostNum = append(mostNum, root.Val)
		}
		preNode = root
		dfs(root.Right)
	}
	dfs(root)
	return mostNum
}

func TestT501(t *testing.T) {
	assert.Equal(t, []int{2}, findMode(buildTree([]int{1, null, 2, 2})))
	assert.Equal(t, []int{0}, findMode(buildTree([]int{0})))
	assert.Equal(t, []int{2, 1}, findMode(buildTree([]int{1, null, 2})))
	assert.Equal(t, []int{0, 1}, findMode(buildTree([]int{1, 0, 1, 0, 0, 1, 1, 0})))

}
