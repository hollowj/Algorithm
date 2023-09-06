package stack_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func maxSlidingWindow(nums []int, k int) []int {
//		result := make([]int, 0)
//		for i := 0; i <= len(nums)-k; i++ {
//			result = append(result, max(nums[i:i+k]))
//		}
//		return result
//	}
//
//	func max(arr []int) int {
//		iMax := math.MinInt
//		for _, num := range arr {
//			if num > iMax {
//				iMax = num
//			}
//		}
//		return iMax
//	}
type SingleQueue struct {
	queue []int
	nums  []int
	k     int
}

func NewSingleQueue(nums []int, k int) *SingleQueue {
	return &SingleQueue{queue: make([]int, 0), nums: nums, k: k}
}

func (self *SingleQueue) push(index int) {
	if !self.empty() && index-self.peekFront() >= self.k {
		self.popFront()
	}
	for !self.empty() && self.nums[self.peekEnd()] <= self.nums[index] {
		self.popEnd()
	}
	self.queue = append(self.queue, index)
}
func (self *SingleQueue) peekFront() int {
	value := self.queue[0]
	return value
}
func (self *SingleQueue) popFront() int {
	value := self.peekFront()
	self.queue = self.queue[1:]
	return value
}
func (self *SingleQueue) peekEnd() int {
	size := len(self.queue)
	value := self.queue[size-1]
	return value
}
func (self *SingleQueue) popEnd() int {
	value := self.peekEnd()
	self.queue = self.queue[:len(self.queue)-1]
	return value
}
func (self *SingleQueue) empty() bool {

	return len(self.queue) == 0
}
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewSingleQueue(nums, k)
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		queue.push(i)
		if i >= k-1 {
			iMaxIndex := queue.peekFront()
			result = append(result, nums[iMaxIndex])
		}
	}
	return result
}

func TestT239_test(t *testing.T) {
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	assert.Equal(t, []int{1}, maxSlidingWindow([]int{1}, 1))
	assert.Equal(t, []int{1, -1}, maxSlidingWindow([]int{1, -1}, 1))
	assert.Equal(t, []int{3, 3, 2, 5}, maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
