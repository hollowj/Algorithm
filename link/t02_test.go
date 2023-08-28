package link

import (
	"testing"
)

//
//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	curA := headA
//	curB := headB
//	sizeA := 0
//	sizeB := 0
//	for curA != nil || curB != nil {
//		if curA != nil {
//			curA = curA.Next
//			sizeA++
//		}
//		if curB != nil {
//			curB = curB.Next
//			sizeB++
//		}
//	}
//	var longHead *ListNode
//	var shortHead *ListNode
//	var delta int
//	if sizeA > sizeB {
//		longHead = headA
//		shortHead = headB
//		delta = sizeA - sizeB
//	} else {
//		longHead = headB
//		shortHead = headA
//		delta = sizeB - sizeA
//	}
//	for i := 0; i < delta; i++ {
//		longHead = longHead.Next
//	}
//	for longHead != nil && shortHead != nil {
//		if longHead == shortHead {
//			return shortHead
//		}
//		longHead = longHead.Next
//		shortHead = shortHead.Next
//	}
//	return nil
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	sizeA := 0
	sizeB := 0
	for curA != nil || curB != nil {
		if curA != nil {
			curA = curA.Next
			sizeA++
		}
		if curB != nil {
			curB = curB.Next
			sizeB++
		}
	}
	if sizeA < sizeB {
		headA, headB = headB, headA
		sizeA, sizeB = sizeB, sizeA
	}
	delta := sizeA - sizeB

	for i := 0; i < delta; i++ {
		headA = headA.Next
	}
	for headA != nil && headB != nil {
		if headA == headB {
			return headB
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
}

func TestT02(t *testing.T) {
}
