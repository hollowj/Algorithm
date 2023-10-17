package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type pos struct {
	x, y int
}

func largestIsland(grid [][]int) int {
	findSet := make(map[pos]*pos)
	result := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		result[i] = make([]int, len(grid[0]))
	}
	var dfs func(x, y int) int
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	getRootPos := func(x, y int) *pos {
		cur := pos{x, y}
		for findSet[cur] != nil {
			tmp := *findSet[cur]
			if cur == tmp {
				break
			}
			cur = tmp
		}
		return &cur
	}
	dfs = func(x, y int) int {
		grid[x][y] = 2
		count := 1
		for _, dir := range dirs {
			nextX, nextY := x+dir[0], y+dir[1]
			if nextX < 0 || nextY < 0 || nextY >= len(grid[0]) || nextX >= len(grid) {
				continue
			}
			if grid[nextX][nextY] == 1 {
				findSet[pos{nextX, nextY}] = getRootPos(x, y)
				count += dfs(nextX, nextY)
			}

		}
		return count
	}
	getCount := func(x, y int) int {
		rootPos := getRootPos(x, y)
		return result[rootPos.x][rootPos.y]
	}
	iMaxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				result[i][j] = dfs(i, j)
				if result[i][j] > iMaxArea {
					iMaxArea = result[i][j]
				}
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				m := make(map[pos]struct{})
				for _, dir := range dirs {
					nextX, nextY := i+dir[0], j+dir[1]
					if nextX < 0 || nextY < 0 || nextY >= len(grid[0]) || nextX >= len(grid) {
						continue
					}
					if grid[nextX][nextY] == 2 {
						m[*getRootPos(nextX, nextY)] = struct{}{}
					}

				}
				iCount := 1
				for p, _ := range m {
					iCount += getCount(p.x, p.y)
				}
				if iCount > iMaxArea {
					iMaxArea = iCount
				}
			}
		}
	}

	return iMaxArea
}

func TestT827(t *testing.T) {
	assert.Equal(t, 3, largestIsland([][]int{{1, 0}, {0, 1}}))
	assert.Equal(t, 4, largestIsland([][]int{{1, 1}, {1, 0}}))
	assert.Equal(t, 4, largestIsland([][]int{{1, 1}, {1, 1}}))
	//assert.Equal(t, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, pacificAtlantic([][]int{{3, 3, 3, 3, 3, 3}, {3, 0, 3, 3, 0, 3}, {3, 3, 3, 3, 3, 3}}))

}
