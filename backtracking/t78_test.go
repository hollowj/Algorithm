package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking78 struct {
	results [][]int
	result  []int
}

func NewBackTracking78() *BackTracking78 {
	return &BackTracking78{results: make([][]int, 0), result: make([]int, 0)}
}
func (b *BackTracking78) dfs(s []int, begin int) {
	tmp := make([]int, len(b.result))
	copy(tmp, b.result)
	b.results = append(b.results, tmp)
	if begin == len(s) {
		return
	}
	for i := begin; i < len(s); i++ {
		b.result = append(b.result, s[i])
		b.dfs(s, i+1)
		b.result = b.result[:len(b.result)-1]
	}
}
func subsets(nums []int) [][]int {
	tracking131 := NewBackTracking78()
	tracking131.dfs(nums, 0)
	return tracking131.results
}
func TestT78(t *testing.T) {
	assert.Equal(t, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}, subsets([]int{1, 2, 3}))
	assert.Equal(t, [][]int{{}, {0}}, subsets([]int{0}))
}
