package backtracking

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking40 struct {
	results [][]int
	result  []int
	used    []bool
}

func NewBackTracking40() *BackTracking40 {
	return &BackTracking40{results: make([][]int, 0), result: make([]int, 0)}
}
func (b *BackTracking40) dfs(candidates []int, target int, begin int, iSum int) {
	if iSum > target {
		return
	}
	if iSum == target {
		tmp := make([]int, len(b.result))
		copy(tmp, b.result)
		b.results = append(b.results, tmp)
		return
	}
	for i := begin; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] && b.used[i-1] == false {
			continue
		}
		iSum += candidates[i]
		b.result = append(b.result, candidates[i])
		b.used[i] = true
		b.dfs(candidates, target, i+1, iSum)
		b.used[i] = false
		iSum -= candidates[i]
		b.result = b.result[:len(b.result)-1]
	}
}
func combinationSum2(candidates []int, target int) [][]int {
	b := NewBackTracking40()
	b.used = make([]bool, len(candidates))
	sort.Ints(candidates)
	b.dfs(candidates, target, 0, 0)
	return b.results
}
func TestT40(t *testing.T) {
	assert.Equal(t, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	assert.Equal(t, [][]int{{1, 2, 2}, {5}}, combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}
