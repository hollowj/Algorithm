package link

type Node struct {
	next *Node
	val  int
}

func NewNode(val int) *Node {
	return &Node{val: val}
}

type SortedList interface {
	Add(num int)
	Search(num int) bool
	ToArray() []int
	Erase(num int) bool
	Range(begin, end int) []int
}

type SortedLinkList struct {
	head *Node
}

func NewSortedLinkList() *SortedLinkList {
	return &SortedLinkList{head: NewNode(0)}
}

func (l *SortedLinkList) Add(num int) {
	node := NewNode(num)
	tmp := l.head
	for tmp.next != nil {
		if num < tmp.next.val {
			break
		}
		tmp = tmp.next
	}
	if tmp.next != nil {
		node.next = tmp.next
	}
	tmp.next = node

}
func (l *SortedLinkList) ToArray() []int {
	tmp := l.head
	nums := make([]int, 0)
	for tmp.next != nil {
		nums = append(nums, tmp.next.val)
		tmp = tmp.next
	}
	return nums
}
func (l *SortedLinkList) Search(num int) bool {
	tmp := l.head
	for tmp.next != nil {
		if tmp.next.val == num {
			return true
		}
		tmp = tmp.next
	}
	return false
}
func (l *SortedLinkList) Erase(num int) bool {
	tmp := l.head
	for tmp.next != nil {
		if tmp.next.val == num {
			tmp.next = tmp.next.next
			return true
		}
		tmp = tmp.next
	}
	return false
}
func (l *SortedLinkList) Range(begin, end int) []int {
	nums := make([]int, 0)

	tmp := l.head
	for tmp.next != nil {
		if tmp.next.val >= begin && tmp.next.val <= end {
			nums = append(nums, tmp.next.val)
		}
		tmp = tmp.next
	}
	return nums
}
