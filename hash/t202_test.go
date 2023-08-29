package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func isHappy(n int) bool {
//		m := make(map[int]struct{})
//		return isHappy_1(n, m)
//	}
//
//	func isHappy_1(n int, m map[int]struct{}) bool {
//		if _, ok := m[n]; ok {
//			return false
//		}
//		m[n] = struct{}{}
//		sum := 0
//		for n != 0 {
//			x := n % 10
//			sum += x * x
//			n = n / 10
//		}
//		if sum == 1 {
//			return true
//		}
//		return isHappy_1(sum, m)
//	}
//
//	func isHappy(n int) bool {
//		m := make(map[int]struct{})
//		return isHappy_1(n, m)
//	}
func getSum(n int) int {
	sum := 0
	for n != 0 {
		x := n % 10
		sum += x * x
		n = n / 10
	}
	return sum
}
func isHappy(n int) bool {
	m := make(map[int]struct{})
	for true {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
		sum := getSum(n)
		if sum == 1 {
			return true
		}
		n = sum
	}
	return false
}
func TestT202(t *testing.T) {
	assert.Equal(t, true, isHappy(19))
	assert.Equal(t, false, isHappy(2))
}
