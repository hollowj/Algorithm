package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//func allPathsSourceTarget(graph [][]int) [][]int {
//	result := [][]int{{0}}
//	i := 0
//	for i < len(graph)-1 {
//		tmp := make([][]int, 0)
//		for _, track := range result {
//			curNum := track[len(track)-1]
//			if curNum == len(graph)-1 {
//				tmp = append(tmp, track)
//			} else {
//				canMovePoints := graph[curNum]
//				for _, point := range canMovePoints {
//					newTrack := append([]int{}, track...)
//					newTrack = append(newTrack, point)
//					tmp = append(tmp, newTrack)
//				}
//			}
//
//		}
//		result = tmp
//
//		i++
//	}
//
//	return result
//}

func allPathsSourceTarget(graph [][]int) [][]int {
	result := [][]int{}
	var dfs func(graph [][]int, visited []int, pos int)
	dfs = func(graph [][]int, visited []int, pos int) {
		visited = append(visited, pos)
		if pos == len(graph)-1 {
			nums := append([]int{}, visited...)
			result = append(result, nums)
			return
		}
		canMovePoints := graph[pos]
		for _, point := range canMovePoints {
			dfs(graph, visited, point)
		}
	}
	dfs(graph, nil, 0)
	return result
}

func TestT797(t *testing.T) {
	assert.Equal(t, [][]int{{0, 1, 3}, {0, 2, 3}}, allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	assert.Equal(t, [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}, {0, 1, 4}}, allPathsSourceTarget([][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}))
}

func TestSlice(t *testing.T) {
	nums := make([]int, 0, 5)
	fmt.Printf("%p %v\n", nums, nums)
	for i := 0; i < 6; i++ {
		nums = append(nums, i)
		fmt.Printf("%d %p %v\n", i, nums, nums)
		if i == 3 {
			n := nums[:2]
			num := i * 10
			n = append(n, num)
			fmt.Printf("%d %p %v\n", num, nums, nums)
		}
	}
	fmt.Printf("%p %v\n", nums, nums)
	nums = append([]int{}, nums...)
	fmt.Printf("%p %v\n", nums, nums)
}
