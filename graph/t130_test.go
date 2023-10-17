package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func solve(board [][]byte) {
	var dfs func(x, y int)
	dirs := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	dfs = func(x, y int) {
		board[x][y] = 'Y'
		for _, dir := range dirs {
			nextX, nextY := x+dir[0], y+dir[1]
			if nextX < 0 || nextY < 0 || nextY >= len(board[0]) || nextX >= len(board) {
				continue
			}
			if board[nextX][nextY] == 'X' || board[nextX][nextY] == 'Y' {
				continue
			}
			dfs(nextX, nextY)

		}

	}
	for i := 0; i < len(board); i++ {
		x, y := i, 0
		if board[x][y] == 'O' {
			dfs(x, y)
		}
		x, y = i, len(board[0])-1
		if board[x][y] == 'O' {
			dfs(x, y)
		}
	}
	for j := 1; j < len(board[0]); j++ {
		x, y := 0, j
		if board[x][y] == 'O' {
			dfs(x, y)
		}
		x, y = len(board)-1, j
		if board[x][y] == 'O' {
			dfs(x, y)
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'Y' {
				board[i][j] = 'O'
			}
		}
	}
}

func TestT130(t *testing.T) {
	data := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(data)
	assert.Equal(t, [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}, data)
	data = [][]byte{{'X'}}
	solve(data)
	assert.Equal(t, [][]byte{{'X'}}, data)
	data = [][]byte{{'O', 'O'}, {'O', 'O'}}
	solve(data)
	assert.Equal(t, [][]byte{{'O', 'O'}, {'O', 'O'}}, data)
	data = [][]byte{
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X', 'X', 'O', 'X', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'O', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'X', 'O', 'X', 'X', 'O', 'O', 'X', 'X', 'X'},
		{'O', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'O', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X', 'X', 'O', 'X', 'X', 'O', 'O'},
		{'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'}}
	solve(data)
	assert.Equal(t, [][]byte{
		{'X', 'O', 'O', 'X', 'X', 'X', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X', 'X', 'O', 'O'},
		{'X', 'X', 'X', 'O', 'O', 'X', 'O', 'X', 'X', 'O'}}, data)
	data = [][]byte{
		{'O', 'X', 'O', 'O', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X'},
		{'O', 'O', 'X', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'O', 'O', 'O'}}
	solve(data)
	assert.Equal(t, [][]byte{
		{'O', 'X', 'O', 'O', 'X', 'X'},
		{'O', 'X', 'X', 'X', 'O', 'X'},
		{'X', 'O', 'O', 'X', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'X', 'X'},
		{'O', 'O', 'X', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'O', 'O', 'O'}}, data)
}
