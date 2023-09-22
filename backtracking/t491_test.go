package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking491 struct {
	results [][]int
	result  []int
}

func NewBackTracking491() *BackTracking491 {
	return &BackTracking491{results: make([][]int, 0), result: make([]int, 0)}
}
func (b *BackTracking491) dfs(s []int, begin int) {
	if len(b.result) >= 2 {
		tmp := make([]int, len(b.result))
		copy(tmp, b.result)
		b.results = append(b.results, tmp)
	}
	if begin == len(s) {
		return
	}
	m := make(map[int]struct{})
	for i := begin; i < len(s); i++ {
		_, ok := m[s[i]]
		if i > begin && s[i] == s[i-1] || ok {
			continue
		}
		if len(b.result) == 0 || s[i] >= b.result[len(b.result)-1] {
			b.result = append(b.result, s[i])
			m[s[i]] = struct{}{}
			b.dfs(s, i+1)
			b.result = b.result[:len(b.result)-1]
		}
	}
}
func findSubsequences(nums []int) [][]int {
	tracking491 := NewBackTracking491()
	tracking491.dfs(nums, 0)
	return tracking491.results
}
func TestT491(t *testing.T) {
	assert.Equal(t, [][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}}, findSubsequences([]int{4, 6, 7, 7}))
	assert.Equal(t, [][]int{{4, 4}}, findSubsequences([]int{4, 4, 3, 2, 1}))
	assert.Equal(t, [][]int{{4, 4}}, findSubsequences([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 1, 1, 1, 1}))
}
