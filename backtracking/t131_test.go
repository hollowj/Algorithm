package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking131 struct {
	results [][]string
	result  []string
	used    []bool
}

func NewBackTracking131() *BackTracking131 {
	return &BackTracking131{results: make([][]string, 0), result: make([]string, 0)}
}
func (b *BackTracking131) dfs(s string, begin int) {
	if begin == len(s) {
		tmp := make([]string, len(b.result))
		copy(tmp, b.result)
		b.results = append(b.results, tmp)
		return
	}
	for i := begin; i < len(s); i++ {
		if b.isHuiWen(s[begin : i+1]) {
			b.result = append(b.result, s[begin:i+1])
			b.dfs(s, i+1)
			b.result = b.result[:len(b.result)-1]
		}
	}
}
func (b *BackTracking131) isHuiWen(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func partition(s string) [][]string {
	tracking131 := NewBackTracking131()
	tracking131.dfs(s, 0)
	return tracking131.results
}
func TestT131(t *testing.T) {
	assert.Equal(t, [][]string{{"a", "a", "b"}, {"aa", "b"}}, partition("aab"))
	assert.Equal(t, [][]string{{"a"}}, partition("a"))
}
