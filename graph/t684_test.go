package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findRedundantConnection(edges [][]int) []int {
	unionSet := NewUnionSet(len(edges) + 1)
	for _, edge := range edges {
		from := edge[0]
		to := edge[1]
		if unionSet.IsSame(from, to) {
			return edge
		}
		unionSet.Join(from, to)
	}
	return []int{}
}

func TestT684(t *testing.T) {
	assert.Equal(t, []int{2, 3}, findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	assert.Equal(t, []int{1, 4}, findRedundantConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}))
}
