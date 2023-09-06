package stack_queue

import (
	"testing"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/stretchr/testify/assert"
)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	queue := priorityqueue.NewWith(func(a, b interface{}) int {
		return a.([]int)[1] - b.([]int)[1]
	})
	for key, v := range m {
		if queue.Size() >= k {
			value, _ := queue.Peek()
			if value.([]int)[1] > v {
				continue
			}
			queue.Dequeue()
		}
		queue.Enqueue([]int{key, v})
	}
	result := make([]int, 0)
	for !queue.Empty() {
		value, _ := queue.Dequeue()
		result = append(result, value.([]int)[0])
	}
	return result
}
func TestT347(t *testing.T) {
	assert.Equal(t, []int{1, 2}, topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	assert.Equal(t, []int{1}, topKFrequent([]int{1}, 1))
}
