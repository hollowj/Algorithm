package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
//func numIslands(grid [][]byte) int {
//	var dfs func(x, y int)
//	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
//	dfs = func(x, y int) {
//		grid[x][y] = '2'
//		for _, t := range dirs {
//			x1, y1 := x+t[0], y+t[1]
//			if x1 < 0 || y1 < 0 || y1 >= len(grid[0]) || x1 >= len(grid) {
//				continue
//			}
//			if grid[x1][y1] == '0' || grid[x1][y1] == '2' {
//				continue
//			}
//			dfs(x1, y1)
//		}
//
//	}
//	var count int
//	for i := 0; i < len(grid); i++ {
//		for j := 0; j < len(grid[0]); j++ {
//			if grid[i][j] == '1' {
//				dfs(i, j)
//				count++
//			}
//		}
//	}
//	return count
//}

func numIslands(grid [][]byte) int {
	queue := make([][]int, 0)
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	bfs := func(x, y int) {
		queue = append(queue, []int{x, y})
		grid[x][y] = '2'
		for len(queue) != 0 {
			p := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			x, y = p[0], p[1]
			for _, dir := range dirs {
				x1, y1 := x+dir[0], y+dir[1]
				if x1 < 0 || y1 < 0 || y1 >= len(grid[0]) || x1 >= len(grid) {
					continue
				}
				if grid[x1][y1] == '0' || grid[x1][y1] == '2' {
					continue
				}
				grid[x1][y1] = '2'
				queue = append(queue, []int{x1, y1})
			}
		}
	}
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				bfs(i, j)
				count++
			}
		}
	}
	return count
}
func TestT200(t *testing.T) {
	assert.Equal(t, 1, numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
	assert.Equal(t, 3, numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
	assert.Equal(t, 1, numIslands([][]byte{{'1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1', '1'}}))
}
