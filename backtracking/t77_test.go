package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var results [][]int
var result []int

func combine(n int, k int) [][]int {
	results = make([][]int, 0)
	result = make([]int, 0, k)
	handle(1, n, k)
	return results
}
func handle(begin, n, k int) {
	if len(result) == k {
		res := make([]int, k)
		copy(res, result)
		results = append(results, res)
		return
	}
	for i := begin; i <= n; i++ {
		if len(result)+n-i+1 < k {
			break
		}
		result = append(result, i)
		handle(i+1, n, k)
		result = (result)[:len(result)-1]
	}
}
func TestT77(t *testing.T) {
	assert.Equal(t, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}, combine(4, 2))
	assert.Equal(t, [][]int{{1}}, combine(1, 1))
}
