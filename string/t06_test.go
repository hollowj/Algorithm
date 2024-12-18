package string

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	arr := make([][]byte, numRows)
	down := 1
	right := 2
	dir := down
	m := 0
	for i := 0; i < len(s); i++ {
		if m == numRows {
			dir = right
			m -= 2
		}
		if m == 0 {
			dir = down
		}
		switch dir {
		case down:
			arr[m] = append(arr[m], s[i])
			m++
		case right:
			arr[m] = append(arr[m], s[i])
			m--
		}

	}

	return string(bytes.Join(arr, nil))
}

func TestT06(t *testing.T) {
	assert.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))
	assert.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
	assert.Equal(t, "A", convert("A", 1))
	assert.Equal(t, "AB", convert("AB", 1))
}
