package bak

import (
	"fmt"
	"strings"
	"testing"
)

func initEmptyMap(n int) [][]string {
	emptyMap := make([][]string, n)
	for i := 0; i < n; i++ {
		bareString := make([]string, n)
		for i := 0; i < n; i++ {
			bareString[i] = "."
		}
		emptyMap[i] = bareString
	}
	return emptyMap
}

func solveNQueens(n int) [][]string {
	resultMap := make([][]string, 0)
	emptyMap := initEmptyMap(n)
	solve(&resultMap, emptyMap, 0)
	return resultMap
}
func buildResult(m [][]string) []string {
	result := make([]string, len(m))
	for i, row := range m {
		result[i] = strings.Join(row, "")
	}
	return result
}
func solve(resultMap *[][]string, m [][]string, x int) {
	if x == len(m) {
		*resultMap = append(*resultMap, buildResult(m))
		return
	}
	for i := 0; i < len(m); i++ {
		if checkPos(m, x, i) {
			m[x][i] = "Q"
			solve(resultMap, m, x+1)
			m[x][i] = "."
		}
	}
}
func checkPos(theMap [][]string, x, y int) bool {
	for i := 0; i <= x; i++ {
		if theMap[i][y] == "Q" || (i+y-x >= 0 && i+y-x < len(theMap) && theMap[i][i+y-x] == "Q") || (-i+y+x >= 0 && -i+y+x < len(theMap) && theMap[i][-i+y+x] == "Q") {
			return false
		}
	}
	return true
}

func TestT5(t *testing.T) {
	queens := solveNQueens(1)
	fmt.Println(queens)
	queens = solveNQueens(2)
	fmt.Println(queens)
	queens = solveNQueens(3)
	fmt.Println(queens)
	queens = solveNQueens(4)
	fmt.Println(queens)
}
