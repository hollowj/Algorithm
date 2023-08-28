package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1 2 3 4 5
// slow fast
// 1		1
// 2		3
// 3		5
// 4   	nil
// ```
// slow fast
// 1		1
// 2		3
// 3		5
// 4		nil
// ```
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func TestT876(t *testing.T) {
	assert.Equal(t, []int{3, 4, 5}, middleNode(NewLinkList([]int{1, 2, 3, 4, 5})).Translate2Arr())
	assert.Equal(t, []int{4, 5, 6}, middleNode(NewLinkList([]int{1, 2, 3, 4, 5, 6})).Translate2Arr())
}
