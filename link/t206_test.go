package link

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Auther struct {
	name string
}

func NewAuther(name string) *Auther {
	return &Auther{name: name}
}

type Book struct {
	auther   *Auther
	bookName string
}

func NewBook(bookName, autherName string) *Book {
	return &Book{auther: NewAuther(autherName), bookName: bookName}
}

func TestStruct(t *testing.T) {
	book := NewBook("海底两万里", "海明威")
	fmt.Printf("1 %p %p %s \n", book, book.auther, book.auther.name)
	Say(*book)
}

func Say(book Book) {
	fmt.Printf("2 %p %p %s \n", &book, book.auther, book.auther.name)
}

//	func reverseList(head *ListNode) *ListNode {
//		cur := head
//		var result *ListNode
//		for cur != nil {
//			last := cur.Next
//			cur.Next = result
//			result = cur
//			cur = last
//		}
//		return result
//	}
func reverseList(head *ListNode) *ListNode {
	return reverseList1(nil, head)

}
func reverseList1(pre, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	} else {
		head := reverseList1(cur, cur.Next)
		cur.Next = pre
		return head
	}
}
func TestT206(t *testing.T) {
	assert.Equal(t, []int{5, 4, 3, 2, 1}, reverseList(NewLinkList([]int{1, 2, 3, 4, 5})).Translate2Arr())
	assert.Equal(t, []int{}, reverseList(NewLinkList([]int{})).Translate2Arr())
}
