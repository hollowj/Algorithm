package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

//	func fib(n int) int {
//		if n < 2 {
//			return n
//		}
//		return fib(n-1) + fib(n-2)
//	}
func TestT509(t *testing.T) {
	assert.Equal(t, 1, fib(2))
	assert.Equal(t, 2, fib(3))
	assert.Equal(t, 3, fib(4))
}
