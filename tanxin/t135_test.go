package tanxin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func candy(ratings []int) int {
	candyCount := make([]int, len(ratings))
	candyCount[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candyCount[i] = candyCount[i-1] + 1
		} else {
			candyCount[i] = 1
		}
	}
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyCount[i] = max(candyCount[i], candyCount[i+1]+1)
		}
	}

	iSum := 0
	for _, count := range candyCount {
		iSum += count
	}
	return iSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestT135(t *testing.T) {
	assert.Equal(t, 5, candy([]int{1, 0, 2}))
	assert.Equal(t, 4, candy([]int{1, 2, 2}))
	assert.Equal(t, 11, candy([]int{1, 3, 4, 5, 2}))
}
