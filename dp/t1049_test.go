package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, num := range stones {
		sum += num
	}

	result := sum / 2
	dp := make([]int, result+1)
	for i := 0; i < len(stones); i++ {
		for j := result; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}

	}
	return sum - 2*dp[result]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func TestT1049(t *testing.T) {
	assert.Equal(t, 1, lastStoneWeightII([]int{2, 7, 1, 8, 1}))
	assert.Equal(t, 5, lastStoneWeightII([]int{31, 26, 33, 21, 40}))
}
