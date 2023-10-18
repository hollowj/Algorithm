package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func islandPerimeter(grid [][]int) int {
//		dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
//		result := 0
//		for i := 0; i < len(grid); i++ {
//			for j := 0; j < len(grid[0]); j++ {
//				if grid[i][j] == 1 {
//					for _, dir := range dirs {
//						nextX, nextY := i+dir[0], j+dir[1]
//						if nextX < 0 || nextY < 0 || nextY >= len(grid[0]) || nextX >= len(grid) {
//							result++
//							continue
//						}
//						if grid[nextX][nextY] == 0 {
//							result++
//						}
//
//					}
//				}
//			}
//		}
//		return result
//	}
func islandPerimeter(grid [][]int) int {
	dirs := [][]int{{-1, 0}, {0, 1}}
	iBlockCount := 0
	iCover := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				iBlockCount++
				for _, dir := range dirs {
					nextX, nextY := i+dir[0], j+dir[1]
					if nextX < 0 || nextY < 0 || nextY >= len(grid[0]) || nextX >= len(grid) {
						continue
					}
					if grid[nextX][nextY] == 1 {
						iCover++
					}
				}
			}
		}
	}
	return iBlockCount*4 - 2*iCover
}
func TestT463(t *testing.T) {
	assert.Equal(t, 16, islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	assert.Equal(t, 4, islandPerimeter([][]int{{1}}))
	assert.Equal(t, 4, islandPerimeter([][]int{{1, 0}}))
	assert.Equal(t, 8, islandPerimeter([][]int{{1, 1}, {1, 1}}))
}
