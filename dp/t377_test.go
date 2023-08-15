package dp

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 1; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}

func TestT377(t *testing.T) {
	assert.Equal(t, 7, combinationSum4([]int{1, 2, 3}, 4))
	assert.Equal(t, 0, combinationSum4([]int{9}, 3))
}

func SliceSum(arr []int64) int64 {
	var sum int64

	for _, val := range arr {
		if val%100000 != 0 {
			sum += val
		}
	}

	return sum
}

// slice_sum_test.go
func FuzzSliceSum(f *testing.F) {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())
	// 语料
	f.Add(10, "abc")

	f.Fuzz(func(t *testing.T, n int, s string) {
		fmt.Println("fuzz-", n, s)
		n %= 20

		var arr []int64
		var expect int64 // 期望值
		var buf strings.Builder
		buf.WriteString(" ")

		for i := 0; i < n; i++ {
			val := rand.Int63() % 1000000
			arr = append(arr, val)
			expect += val
			buf.WriteString(fmt.Sprintf("%d,", val))
		}

		// 自己求和的结果和调用函数求和的结果比对
		assert.Equal(t, expect, SliceSum(arr), buf.String())
	})
}
