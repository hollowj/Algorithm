package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func intersection(nums1 []int, nums2 []int) []int {
	exist := make(map[int]bool)
	for _, num := range nums1 {
		exist[num] = true
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		ok := exist[num]
		if ok {
			exist[num] = false
			result = append(result, num)
		}
	}
	return result
}
func TestT349(t *testing.T) {
	assert.Equal(t, []int{2}, intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	assert.Equal(t, []int{9, 4}, intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
