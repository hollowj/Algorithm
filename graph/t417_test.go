package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func pacificAtlantic(heights [][]int) [][]int {
	results := make([][]int, 0)
	var dfs func(ctn [][]bool, x, y int)
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	pacific := make([][]bool, len(heights))
	atlantic := make([][]bool, len(heights))
	for i := 0; i < len(heights); i++ {
		pacific[i] = make([]bool, len(heights[0]))
		atlantic[i] = make([]bool, len(heights[0]))
	}
	dfs = func(ctn [][]bool, x, y int) {

		ctn[x][y] = true
		for _, dir := range dirs {
			nextX, nextY := x+dir[0], y+dir[1]
			if nextX < 0 || nextY < 0 || nextY >= len(heights[0]) || nextX >= len(heights) {
				continue
			}
			if heights[nextX][nextY] >= heights[x][y] && !ctn[nextX][nextY] {
				dfs(ctn, nextX, nextY)
			}

		}

	}
	for i := 0; i < len(heights); i++ {
		dfs(pacific, i, 0)
		dfs(atlantic, i, len(heights[0])-1)

	}
	for j := 0; j < len(heights[0]); j++ {
		dfs(pacific, 0, j)
		dfs(atlantic, len(heights)-1, j)
	}
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[0]); j++ {
			if pacific[i][j] && atlantic[i][j] {
				results = append(results, []int{i, j})
			}
		}
	}
	return results
}

func TestT417(t *testing.T) {
	assert.Equal(t, [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}, pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}))
	assert.Equal(t, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, pacificAtlantic([][]int{{2, 1}, {1, 2}}))
	//assert.Equal(t, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}, pacificAtlantic([][]int{{3, 3, 3, 3, 3, 3}, {3, 0, 3, 3, 0, 3}, {3, 3, 3, 3, 3, 3}}))

}
