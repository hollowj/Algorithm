package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func swapPairs(head *ListNode) *ListNode {
	cur := head
	var root = head
	var prev *ListNode
	for cur != nil {
		var next2 *ListNode
		next := cur.Next
		if next != nil {
			next2 = next.Next
			next.Next = cur
		} else {
			next = cur
		}
		if root == head {
			root = next
		}
		if prev != nil {
			prev.Next = next
		}
		prev = cur
		cur = next2
	}
	if prev != nil {
		prev.Next = nil
	}
	return root
}
func TestT24(t *testing.T) {
	assert.Equal(t, []int{2, 1, 4, 3}, swapPairs(NewLinkList([]int{1, 2, 3, 4})).Translate2Arr())
	assert.Equal(t, []int{}, swapPairs(NewLinkList([]int{})).Translate2Arr())
	assert.Equal(t, []int{1}, swapPairs(NewLinkList([]int{1})).Translate2Arr())
	assert.Equal(t, []int{2, 1, 3}, swapPairs(NewLinkList([]int{1, 2, 3})).Translate2Arr())
}
