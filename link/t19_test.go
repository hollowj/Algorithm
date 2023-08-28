package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//	func removeNthFromEnd(head *ListNode, n int) *ListNode {
//		var cur = 0
//		return removeNthFromEnd_1(nil, head, n, &cur)
//	}
//
//	func removeNthFromEnd_1(prev, head *ListNode, n int, cur *int) *ListNode {
//		if head == nil {
//			return nil
//		}
//		removeNthFromEnd_1(head, head.Next, n, cur)
//		*cur += 1
//		if *cur == n {
//			if prev != nil {
//				prev.Next = head.Next
//			} else {
//				return head.Next
//			}
//		}
//		if prev == nil {
//			return head
//		}
//		return nil
//	}

//1 2 3 4 5 -----2
//fast=3
//slow fast
//1 	4
//2	5
//3	nil

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := NewListNode(-1)
	dummyNode.Next = head
	slow := dummyNode
	fast := dummyNode
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyNode.Next
}

func TestT19(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 5}, removeNthFromEnd(NewLinkList([]int{1, 2, 3, 4, 5}), 2).Translate2Arr())
	assert.Equal(t, []int{}, removeNthFromEnd(NewLinkList([]int{1}), 1).Translate2Arr())
	assert.Equal(t, []int{1}, removeNthFromEnd(NewLinkList([]int{1, 2}), 1).Translate2Arr())
}
