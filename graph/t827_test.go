package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func largestIsland(grid [][]int) int {
	result := make(map[int]int)

	var dfs func(x, y, iNo int) int
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	m, n := len(grid), len(grid[0])
	dfs = func(x, y, iNo int) int {
		grid[x][y] = iNo
		count := 1
		for _, dir := range dirs {
			nextX, nextY := x+dir[0], y+dir[1]
			if nextX < 0 || nextY < 0 || nextY >= n || nextX >= m {
				continue
			}
			if grid[nextX][nextY] == 1 {
				count += dfs(nextX, nextY, iNo)
			}

		}
		return count
	}
	iNo := 2
	iMaxArea := 0
	isAllGrid := true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				result[iNo] = dfs(i, j, iNo)
				if result[iNo] > iMaxArea {
					iMaxArea = result[iNo]
				}
				iNo++
			} else if grid[i][j] == 0 {
				isAllGrid = false
			}
		}
	}
	if isAllGrid {
		return m * n
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				iCount := 1
				set := make(map[int]struct{})
				for _, dir := range dirs {
					nextX, nextY := i+dir[0], j+dir[1]
					if nextX < 0 || nextY < 0 || nextY >= n || nextX >= m {
						continue
					}
					if grid[nextX][nextY] > 1 {
						if _, ok := set[grid[nextX][nextY]]; !ok {
							iCount += result[grid[nextX][nextY]]
							set[grid[nextX][nextY]] = struct{}{}
						}
					}

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
