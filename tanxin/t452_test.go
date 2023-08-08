package tanxin

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	destroyedCount := 0
	iMinTime := 0
	for destroyedCount < len(points) {
		p := points[destroyedCount]
		destroyedCount++
		for i := destroyedCount; i < len(points); i++ {
			if p[1] >= points[i][0] && p[1] <= points[i][1] {
				destroyedCount++
			} else {
				break
			}
		}
		iMinTime++
	}
	return iMinTime
}
func TestT452(t *testing.T) {
	assert.Equal(t, 2, findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	assert.Equal(t, 4, findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	assert.Equal(t, 2, findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
}
