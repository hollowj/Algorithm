package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
func TestT141(t *testing.T) {
	assert.Equal(t, true, hasCycle(NewLinkList([]int{3, 2, 0, -4})))
	assert.Equal(t, true, hasCycle(NewLinkList([]int{1, 2})))
	assert.Equal(t, false, hasCycle(NewLinkList([]int{1})))
}
