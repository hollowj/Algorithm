package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type BackTracking37 struct {
}

func NewBackTracking37() *BackTracking37 {
	return &BackTracking37{}
}
func (b *BackTracking37) dfs(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for i := '1'; i <= '9'; i++ {
					ch := byte(i)
					if b.checkValid(board, row, col, ch) {
						board[row][col] = ch
						if b.dfs(board) {
							return true
						}
						board[row][col] = '.'
					}
				}
				return false
			}
		}
	}

	return true
}
func (b *BackTracking37) checkValid(board [][]byte, row, col int, ch byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == ch {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if board[i][col] == ch {
			return false
		}
	}
	startRow := row / 3 * 3
	startCol := col / 3 * 3
	for row9 := startRow; row9 < startRow+3; row9++ {
		for col9 := startCol; col9 < startCol+3; col9++ {
			if board[row9][col9] == ch {
				return false
			}
		}
	}
	return true
}
func solveSudoku(board [][]byte) {
	tracking491 := NewBackTracking37()
	tracking491.dfs(board)
}
func TestT37(t *testing.T) {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	assert.Equal(t, [][]byte{{'5', '3', '4', '6', '7', '8', '9', '1', '2'}, {'6', '7', '2', '1', '9', '5', '3', '4', '8'}, {'1', '9', '8', '3', '4', '2', '5', '6', '7'}, {'8', '5', '9', '7', '6', '1', '4', '2', '3'}, {'4', '2', '6', '8', '5', '3', '7', '9', '1'}, {'7', '1', '3', '9', '2', '4', '8', '5', '6'}, {'9', '6', '1', '5', '3', '7', '2', '8', '4'}, {'2', '8', '7', '4', '1', '9', '6', '3', '5'}, {'3', '4', '5', '2', '8', '6', '1', '7', '9'}}, board)
}
