package backtracking

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking47 struct {
	results [][]int
	result  []int
}

func NewBackTracking47() *BackTracking47 {
	return &BackTracking47{results: make([][]int, 0), result: make([]int, 0)}
}
func (b *BackTracking47) dfs(s []int, used map[int]bool) {
	if len(b.result) == len(s) {
		tmp := make([]int, len(b.result))
		copy(tmp, b.result)
		b.results = append(b.results, tmp)
		return
	}
	for i := 0; i < len(s); i++ {
		if used[i] || (i > 0 && s[i-1] == s[i] && used[i-1] == false) {
			continue
		}
		b.result = append(b.result, s[i])
		used[i] = true
		b.dfs(s, used)
		b.result = b.result[:len(b.result)-1]
		used[i] = false

	}
}
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	tracking491 := NewBackTracking47()
	used := make(map[int]bool)
	tracking491.dfs(nums, used)
	return tracking491.results
}
func TestT47(t *testing.T) {
	assert.Equal(t, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}, permuteUnique([]int{1, 1, 2}))
	assert.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permuteUnique([]int{1, 2, 3}))
}
func TestMap(t *testing.T) {
	used := make(map[int]bool)
	fmt.Println(used[1])
}
