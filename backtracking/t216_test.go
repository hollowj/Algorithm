package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	results216 [][]int
	result216  []int
)

func combinationSum3(k int, n int) [][]int {
	results216 = make([][]int, 0)
	result216 = make([]int, 0)
	dfs(1, k, n, 0)
	return results216
}
func dfs(begin, k, n, isum int) {
	if len(result216) == k && isum == n {
		tmp := make([]int, k)
		copy(tmp, result216)
		results216 = append(results216, tmp)
		return
	}
	for i := begin; i <= 9; i++ {
		if isum+i > n || 9-i+1 < k-len(result216) {
			break
		}
		isum += i
		result216 = append(result216, i)
		dfs(i+1, k, n, isum)
		isum -= i
		result216 = result216[:len(result216)-1]
	}

}
func sum(nums []int) int {
	iSum := 0
	for i := 0; i < len(nums); i++ {
		iSum += nums[i]
	}
	return iSum
}
func TestT216(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 4}}, combinationSum3(3, 7))
	assert.Equal(t, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}, combinationSum3(3, 9))
}
