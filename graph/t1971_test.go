package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func validPath(n int, edges [][]int, source int, destination int) bool {
	unionSet := NewUnionSet(n)
	for _, edge := range edges {
		unionSet.Join(edge[0], edge[1])
	}
	return unionSet.IsSame(source, destination)
}

func TestT1971(t *testing.T) {
	assert.Equal(t, true, validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	assert.Equal(t, false, validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}, 0, 5))
}
