package backtracking

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking90 struct {
	results [][]int
	result  []int
}

func NewBackTracking90() *BackTracking90 {
	return &BackTracking90{results: make([][]int, 0), result: make([]int, 0)}
}
func (b *BackTracking90) dfs(s []int, begin int) {
	tmp := make([]int, len(b.result))
	copy(tmp, b.result)
	b.results = append(b.results, tmp)
	if begin == len(s) {
		return
	}
	for i := begin; i < len(s); i++ {
		if i > begin && s[i] == s[i-1] {
			continue
		}
		b.result = append(b.result, s[i])
		b.dfs(s, i+1)
		b.result = b.result[:len(b.result)-1]
	}
}
func subsetsWithDup(nums []int) [][]int {
	tracking131 := NewBackTracking90()
	sort.Ints(nums)
	tracking131.dfs(nums, 0)
	return tracking131.results
}
func TestT90(t *testing.T) {
	assert.Equal(t, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}, subsetsWithDup([]int{1, 2, 2}))
	assert.Equal(t, [][]int{{}, {0}}, subsetsWithDup([]int{0}))
}
