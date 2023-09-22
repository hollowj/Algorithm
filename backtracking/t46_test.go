package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking46 struct {
	results [][]int
	result  []int
}

func NewBackTracking46() *BackTracking46 {
	return &BackTracking46{results: make([][]int, 0), result: make([]int, 0)}
}
func (b *BackTracking46) dfs(s []int, used map[int]bool) {
	if len(b.result) >= len(s) {
		tmp := make([]int, len(b.result))
		copy(tmp, b.result)
		b.results = append(b.results, tmp)
		return
	}

	for i := 0; i < len(s); i++ {
		if used[s[i]] {
			continue
		}
		b.result = append(b.result, s[i])
		used[s[i]] = true
		b.dfs(s, used)
		b.result = b.result[:len(b.result)-1]
		used[s[i]] = false

	}
}
func permute(nums []int) [][]int {
	tracking491 := NewBackTracking46()
	used := make(map[int]bool)
	tracking491.dfs(nums, used)
	return tracking491.results
}
func TestT46(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute([]int{1, 2, 3}))
	assert.Equal(t, [][]int{{0, 1}, {1, 0}}, permute([]int{0, 1}))
	assert.Equal(t, [][]int{{1}}, permute([]int{1}))
}
