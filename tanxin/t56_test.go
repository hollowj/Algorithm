package tanxin

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for index := 1; index < len(intervals); index++ {
		t := intervals[index]
		lastTuple := result[len(result)-1]
		if t[0] > lastTuple[1] {
			result = append(result, t)
		} else {
			result[len(result)-1][1] = max(lastTuple[1], t[1])
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestT56(t *testing.T) {
	assert.Equal(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	assert.Equal(t, [][]int{{1, 5}}, merge([][]int{{1, 4}, {4, 5}}))
	assert.Equal(t, [][]int{{1, 10}}, merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
