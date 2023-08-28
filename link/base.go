package link

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}
func (node *ListNode) Translate2Arr() []int {
	if node == nil {
		return []int{}
	}
	arr := make([]int, 0)
	for node != nil {
		arr = append(arr, node.Val)
		node = node.Next
	}
	return arr
}

func NewLinkList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	root := NewListNode(arr[0])
	t := root
	for i := 1; i < len(arr); i++ {
		t.Next = NewListNode(arr[i])
		t = t.Next
	}
	return root
}
