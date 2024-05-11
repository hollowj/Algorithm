package link

import (
	"math/rand"
)

type SkipNode struct {
	next []*SkipNode
	val  int
}

func NewSkipNode(val, level int) *SkipNode {
	return &SkipNode{val: val, next: make([]*SkipNode, level+1)}
}

type SkipList struct {
	head *SkipNode
}

func NewSkipList() *SkipList {
	return &SkipList{head: NewSkipNode(0, 0)}
}

func (l *SkipList) Roll() int {
	var level int
	for rand.Intn(2) == 0 {
		level++
	}
	return level
}
func (l *SkipList) Add(num int) {
	level := l.Roll()
	for len(l.head.next) < level+1 {
		l.head.next = append(l.head.next, nil)
	}
	node := NewSkipNode(num, level)
	cur := l.head
	for i := level; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val <= num {
			cur = cur.next[i]
		}
		node.next[i] = cur.next[i]
		cur.next[i] = node
	}

}
func (l *SkipList) ToArray() []int {
	level := len(l.head.next) - 1
	tmp := l.head
	nums := make([]int, 0)
	for tmp.next[level] != nil {
		nums = append(nums, tmp.next[level].val)
		tmp = tmp.next[level]
	}
	return nums
}
func (l *SkipList) Search(num int) bool {
	cur := l.head
	for i := len(l.head.next) - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		if cur.next[i].val == num {
			return true
		}

	}
	return false
}
func (l *SkipList) Erase(num int) bool {
	ok := false
	cur := l.head
	for i := len(l.head.next) - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		if cur.next[i] != nil && cur.next[i].val == num {
			cur.next[i] = cur.next[i].next[i]
			ok = true
		}

	}
	nilLevelCount := 0
	for i := len(l.head.next) - 1; i >= 0 && l.head.next[i] == nil; i-- {
		nilLevelCount++
	}
	if nilLevelCount > 0 {
		l.head.next = l.head.next[:len(l.head.next)-nilLevelCount]
	}
	return ok
}
func (l *SkipList) Range(begin, end int) []int {
	nums := make([]int, 0)
	cur := l.head
	for i := len(l.head.next) - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < begin {
			cur = cur.next[i]
		}
	}
	for cur.next[0] != nil && cur.next[0].val <= end {
		nums = append(nums, cur.next[0].val)
		cur = cur.next[0]
	}
	return nums
}
