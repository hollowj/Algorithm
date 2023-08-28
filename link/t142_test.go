package link

import (
	"testing"
)

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			headTmp := head
			for headTmp != fast {
				fast = fast.Next
				headTmp = headTmp.Next
			}
			return headTmp
		}
	}

	return nil
}

func TestT142(t *testing.T) {

}
