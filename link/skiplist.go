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

type Skiplist struct {
	head *SkipNode
}

func NewSkipList() *Skiplist {
	return &Skiplist{head: NewSkipNode(0, 0)}
}

func (this *Skiplist) Roll() int {
	var level int
	for rand.Intn(2) == 0 {
		level++
	}
	return level
}
func (this *Skiplist) Add(num int) {
	level := this.Roll()
	for len(this.head.next) < level+1 {
		this.head.next = append(this.head.next, nil)
	}
	node := NewSkipNode(num, level)
	cur := this.head
	for i := level; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val <= num {
			cur = cur.next[i]
		}
		node.next[i] = cur.next[i]
		cur.next[i] = node
	}

}
func (this *Skiplist) ToArray() []int {
	level := len(this.head.next) - 1
	tmp := this.head
	nums := make([]int, 0)
	for tmp.next[level] != nil {
		nums = append(nums, tmp.next[level].val)
		tmp = tmp.next[level]
	}
	return nums
}
func (this *Skiplist) Search(num int) bool {
	cur := this.head
	for i := len(this.head.next) - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		if cur.next[i].val == num {
			return true
		}

	}
	return false
}
func (this *Skiplist) Erase(num int) bool {
	ok := false
	cur := this.head
	for i := len(this.head.next) - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		if cur.next[i] != nil && cur.next[i].val == num {
			cur.next[i] = cur.next[i].next[i]
			ok = true
		}

	}
	nilLevelCount := 0
	for i := len(this.head.next) - 1; i >= 0 && this.head.next[i] == nil; i-- {
		nilLevelCount++
	}
	if nilLevelCount > 0 {
		this.head.next = this.head.next[:len(this.head.next)-nilLevelCount]
	}
	return ok
}
func (this *Skiplist) Range(begin, end int) []int {
	nums := make([]int, 0)
	cur := this.head
	for i := len(this.head.next) - 1; i >= 0; i-- {
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
