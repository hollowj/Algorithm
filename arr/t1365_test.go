package arr

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func smallerNumbersThanCurrent(nums []int) []int {
//		count := len(nums)
//		target := make([]int, count)
//		for i := 0; i < count; i++ {
//			for j := 0; j < count; j++ {
//				if nums[j] > nums[i] {
//					target[j] += 1
//				}
//			}
//		}
//		return target
//	}
//
//	func smallerNumbersThanCurrent(nums []int) []int {
//		sortedNums := append([]int{}, nums...)
//		sort.Slice(sortedNums, func(i, j int) bool {
//			return sortedNums[i] < sortedNums[j]
//		})
//		m := make(map[int]int)
//		for i := 1; i < len(sortedNums); i++ {
//			j := i
//			for j-1 >= 0 && sortedNums[j-1] == sortedNums[i] {
//				j--
//			}
//			m[sortedNums[i]] = j
//		}
//		target := make([]int, len(nums))
//		for i, num := range nums {
//			target[i] = m[num]
//		}
//		return target
//	}
func smallerNumbersThanCurrent(nums []int) []int {
	count := len(nums)
	sortedNums := make([]int, count)
	copy(sortedNums, nums)
	sort.Ints(sortedNums)
	m := make(map[int]int)
	for i := count - 1; i >= 0; i-- {
		m[sortedNums[i]] = i
	}
	target := make([]int, count)
	for i, num := range nums {
		target[i] = m[num]
	}
	return target
}
func TestT1365(t *testing.T) {
	assert.Equal(t, []int{4, 0, 1, 1, 3}, smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	assert.Equal(t, []int{2, 1, 0, 3}, smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	assert.Equal(t, []int{0, 0, 0, 0}, smallerNumbersThanCurrent([]int{7, 7, 7, 7}))
}
