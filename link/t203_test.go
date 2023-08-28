package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func removeElements(head *ListNode, val int) *ListNode {
//		root := NewListNode(val - 1)
//		root.Next = head
//		slow := root
//		fast := root.Next
//		for fast != nil {
//			if fast.Val == val {
//				slow.Next = fast.Next
//				fast = fast.Next
//			} else {
//				slow = slow.Next
//				fast = fast.Next
//			}
//		}
//		return root.Next
//	}
func removeElements(head *ListNode, val int) *ListNode {
	root := NewListNode(val - 1)
	root.Next = head
	cur := root
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return root.Next
}
func TestT203(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, removeElements(NewLinkList([]int{1, 2, 6, 3, 4, 5, 6}), 6).Translate2Arr())
	assert.Equal(t, []int{}, removeElements(NewLinkList([]int{}), 1).Translate2Arr())
	assert.Equal(t, []int{}, removeElements(NewLinkList([]int{7, 7, 7, 7}), 7).Translate2Arr())
}
