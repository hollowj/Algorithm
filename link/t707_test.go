package link

import (
	"fmt"
	"testing"
)

type MyLinkedList struct {
	root *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{root: NewListNode(0)}
}

func (this *MyLinkedList) Get(index int) int {
	if index > this.size {
		return -1
	}

	i := 0
	cur := this.root.Next
	for cur != nil {
		if i == index {
			return cur.Val
		}
		i++
		cur = cur.Next
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	head := NewListNode(val)
	head.Next = this.root.Next
	this.root.Next = head
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {

	cur := this.root
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = NewListNode(val)

	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	i := 0
	cur := this.root
	for cur != nil {
		if i == index {
			tmp := NewListNode(val)
			tmp.Next = cur.Next
			cur.Next = tmp
			break
		}
		i++
		cur = cur.Next
	}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	i := 0
	cur := this.root
	for cur != nil {
		if i == index {
			cur.Next = cur.Next.Next
			break
		}
		i++
		cur = cur.Next
	}
	this.size--
}

func TestT707(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2)
	fmt.Println(linkedList.Get(1))
	linkedList.DeleteAtIndex(0)
	fmt.Println(linkedList.Get(1))
}
func TestT707_7(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtIndex(0, 10)
	linkedList.AddAtIndex(0, 20)
	linkedList.AddAtIndex(1, 30)
	fmt.Println(linkedList.Get(0))
}
func TestT707_8(t *testing.T) {
	linkedList := Constructor()
	linkedList.AddAtHead(2)
	linkedList.DeleteAtIndex(1)
	linkedList.AddAtHead(2)
	linkedList.AddAtHead(7)
	linkedList.AddAtHead(3)
	linkedList.AddAtHead(2)
	linkedList.AddAtHead(5)
	linkedList.AddAtHead(5)
	fmt.Println(linkedList.Get(5))
	linkedList.DeleteAtIndex(6)
	linkedList.DeleteAtIndex(4)
}
