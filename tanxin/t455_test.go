package tanxin

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort.Slice(g, func(i, j int) bool {
		return g[i] >= g[j]
	})
	sort.Slice(s, func(i, j int) bool {
		return s[i] >= s[j]

	})
	cookieIndex := 0
	for i := 0; i < len(g); i++ {
		if cookieIndex >= len(s) {
			break
		}
		if g[i] <= s[cookieIndex] {
			cookieIndex++
		}
	}
	return cookieIndex
}

func TestT455(t *testing.T) {
	assert.Equal(t, 1, findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	assert.Equal(t, 2, findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
