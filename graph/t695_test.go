package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
//func maxAreaOfIsland(grid [][]int) int {
//	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
//	bfs := func(x, y int) int {
//		queue := [][]int{{x, y}}
//		grid[x][y] = 2
//		iArea := 1
//		for len(queue) != 0 {
//			t := queue[len(queue)-1]
//			queue = queue[:len(queue)-1]
//			x, y = t[0], t[1]
//			for _, dir := range dirs {
//				x1, y1 := x+dir[0], y+dir[1]
//				if x1 < 0 || y1 < 0 || y1 >= len(grid[0]) || x1 >= len(grid) {
//					continue
//				}
//				if grid[x1][y1] == 0 || grid[x1][y1] == 2 {
//					continue
//				}
//				iArea++
//				grid[x1][y1] = 2
//				queue = append(queue, []int{x1, y1})
//			}
//		}
//		return iArea
//
//	}
//	iMaxArea := 0
//	for i := 0; i < len(grid); i++ {
//		for j := 0; j < len(grid[0]); j++ {
//			if grid[i][j] == 1 {
//				iArea := bfs(i, j)
//				if iArea > iMaxArea {
//					iMaxArea = iArea
//				}
//			}
//		}
//	}
//	return iMaxArea
//}

func maxAreaOfIsland(grid [][]int) int {
	var dfs func(x, y int) int
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	dfs = func(x, y int) int {
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
			iArea += dfs(x1, y1)
			grid[x1][y1] = 2
		}

		return iArea

	}
	iMaxArea := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				iArea := dfs(i, j)
				if iArea > iMaxArea {
					iMaxArea = iArea
				}
			}
		}
	}
	return iMaxArea
}
func TestT695(t *testing.T) {
	assert.Equal(t, 6, maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}))
	assert.Equal(t, 0, maxAreaOfIsland([][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
	}))
}
