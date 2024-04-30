package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isTreeAfterRemove(edges [][]int, rmIndex int) bool {
	unionSet := NewUnionSet(len(edges) + 1)
	for index, edge := range edges {
		if index == rmIndex {
			continue
		}
		if unionSet.IsSame(edge[0], edge[1]) {
			return false
		}
		unionSet.Join(edge[1], edge[0])
	}
	return true
}
func findRedundantDirectedConnection(edges [][]int) []int {
	inDegree := make([]int, len(edges)+1)
	needDeleteEdges := make([]int, 0)
	for _, edge := range edges {
		inDegree[edge[1]] += 1

	}
	for point, count := range inDegree {
		if count > 1 {
			for index, edge := range edges {
				if edge[1] == point {
					needDeleteEdges = append(needDeleteEdges, index)
				}
			}
			break
		}
	}
	if len(needDeleteEdges) > 0 {
		for i := len(needDeleteEdges) - 1; i >= 0; i-- {
			rmIndex := needDeleteEdges[i]
			if isTreeAfterRemove(edges, rmIndex) {
				return edges[rmIndex]
			}
		}
	} else {
		unionSet := NewUnionSet(len(edges) + 1)
		for _, edge := range edges {
			from := edge[1]
			to := edge[0]
			if unionSet.IsSame(from, to) {
				return edge
			}
			unionSet.Join(from, to)
		}
	}
	return []int{}
}

func TestT685(t *testing.T) {
	assert.Equal(t, []int{2, 3}, findRedundantDirectedConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
	assert.Equal(t, []int{4, 1}, findRedundantDirectedConnection([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}}))
	assert.Equal(t, []int{2, 1}, findRedundantDirectedConnection([][]int{{2, 1}, {3, 1}, {4, 2}, {1, 4}}))
}
