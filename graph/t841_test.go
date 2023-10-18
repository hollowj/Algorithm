package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canVisitAllRooms(rooms [][]int) bool {
	queue := []int{0}
	visited := make(map[int]struct{})
	visited[0] = struct{}{}
	for len(queue) != 0 {
		iRoomNo := queue[0]
		queue = queue[1:]
		targets := rooms[iRoomNo]
		for _, targetNo := range targets {
			if _, ok := visited[targetNo]; !ok {
				visited[targetNo] = struct{}{}
				queue = append(queue, targetNo)
			}
		}
	}
	return len(visited) == len(rooms)
}
func TestT841(t *testing.T) {
	assert.Equal(t, true, canVisitAllRooms([][]int{{1}, {2}, {3}, {}}))
	assert.Equal(t, false, canVisitAllRooms([][]int{{1, 3}, {3, 0, 1}, {2}, {0}}))
}
