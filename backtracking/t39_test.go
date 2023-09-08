package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var results39 [][]int
var result39 []int

func combinationSum(candidates []int, target int) [][]int {
	results39 = make([][]int, 0)
	result39 = make([]int, 0)
	dfs39(candidates, target, 0, 0)
	return results39
}
func dfs39(candidates []int, target int, begin int, iSum int) {
	if iSum > target {
		return
	}
	if iSum == target {
		tmp := make([]int, len(result39))
		copy(tmp, result39)
		results39 = append(results39, tmp)
		return
	}
	for i := begin; i < len(candidates); i++ {
		iSum += candidates[i]
		result39 = append(result39, candidates[i])
		dfs39(candidates, target, i, iSum)
		iSum -= candidates[i]
		result39 = result39[:len(result39)-1]
	}
}
func TestT39(t *testing.T) {
	assert.Equal(t, [][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
	assert.Equal(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8))
	assert.Equal(t, [][]int{}, combinationSum([]int{2}, 1))
	assert.Equal(t, [][]int{{8, 3}, {7, 4}, {4, 4, 3}}, combinationSum([]int{8, 7, 4, 3}, 11))
}
