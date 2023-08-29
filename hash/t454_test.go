package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			m[num1+num2]++
		}
	}
	count := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			if ct, ok := m[-(num3 + num4)]; ok {
				count += ct
			}
		}
	}
	return count
}

func TestT454(t *testing.T) {
	assert.Equal(t, 2, fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
	assert.Equal(t, 1, fourSumCount([]int{0}, []int{0}, []int{0}, []int{0}))
}
