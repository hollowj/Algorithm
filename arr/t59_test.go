package arr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Direct int

const (
	Direct_UP    Direct = 0
	Direct_Down  Direct = 1
	Direct_Left  Direct = 2
	Direct_Right Direct = 3
)

//	func generateMatrix(n int) [][]int {
//		matrix := make([][]int, n)
//		for i := 0; i < n; i++ {
//			matrix[i] = make([]int, n)
//		}
//		iRow, iCol := 0, -1
//		iDirect := Direct_Right
//		for i := 1; i <= n*n; i++ {
//			switch iDirect {
//			case Direct_UP:
//				iRow--
//				if iRow == 0 || matrix[iRow-1][iCol] != 0 {
//					iDirect = Direct_Right
//				}
//			case Direct_Down:
//				iRow++
//				if iRow == n-1 || matrix[iRow+1][iCol] != 0 {
//					iDirect = Direct_Left
//				}
//			case Direct_Left:
//				iCol--
//				if iCol == 0 || matrix[iRow][iCol-1] != 0 {
//					iDirect = Direct_UP
//				}
//			case Direct_Right:
//				iCol++
//				if iCol == n-1 || matrix[iRow][iCol+1] != 0 {
//					iDirect = Direct_Down
//				}
//
//			}
//			matrix[iRow][iCol] = i
//		}
//		return matrix
//	}
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	loop := n / 2
	startX, startY := 0, 0
	offset := 0
	count := 0
	var iRow, iColumn int
	for loop > 0 {
		iRow, iColumn = startX, startY
		for ; iColumn < n-offset-1; iColumn++ {
			count++
			matrix[startX][iColumn] = count
		}
		for ; iRow < n-offset-1; iRow++ {
			count++
			matrix[iRow][iColumn] = count
		}
		for ; iColumn > offset; iColumn-- {
			count++
			matrix[iRow][iColumn] = count
		}
		for ; iRow > offset; iRow-- {
			count++
			matrix[iRow][iColumn] = count
		}
		startX++
		startY++
		offset++
		loop--
	}
	if n%2 == 1 {
		mid := n / 2
		count++
		matrix[mid][mid] = count
	}
	return matrix
}
func TestT59(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, generateMatrix(3))
	assert.Equal(t, [][]int{{1}}, generateMatrix(1))
	assert.Equal(t, [][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}, generateMatrix(4))
}
