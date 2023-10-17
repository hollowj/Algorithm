package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numEnclaves(grid [][]int) int {
	var dfs func(x, y int) (bool, int)
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	dfs = func(x, y int) (bool, int) {
		isFree := false
		if x == 0 || y == 0 || x == len(grid)-1 || y == len(grid[0])-1 {
			isFree = true
		}
		grid[x][y] = 2
		iArea := 1
		for _, dir := range dirs {
			x1, y1 := x+dir[0], y+dir[1]
			if x1 < 0 || y1 < 0 || y1 >= len(grid[0]) || x1 >= len(grid) {
				continue
			}
			if grid[x1][y1] == 0 || grid[x1][y1] == 2 {
				continue
			}
			ok, count := dfs(x1, y1)
			if ok {
				isFree = true
			}
			iArea += count
		}

		return isFree, iArea

	}
	iAreaSum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				ok, iArea := dfs(i, j)
				if !ok {
					iAreaSum += iArea
				}

			}
		}
	}
	return iAreaSum
}
func TestT1020(t *testing.T) {
	assert.Equal(t, 3, numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
	assert.Equal(t, 0, numEnclaves([][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}))
	assert.Equal(t, 3, numEnclaves([][]int{{0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {1, 1, 0, 0, 0, 1, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1, 0, 1, 0, 0}, {0, 1, 1, 0, 0, 0, 1, 0, 1, 0}, {0, 1, 1, 1, 1, 1, 0, 0, 1, 0}, {0, 0, 1, 0, 1, 1, 1, 1, 0, 1}, {0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, {0, 0, 1, 0, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 0, 1, 1, 0, 0, 0, 0}, {0, 0, 0, 0, 1, 1, 0, 0, 0, 1}}))
}
