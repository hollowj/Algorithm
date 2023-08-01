package bak

import (
	"container/heap"
	"fmt"
	"testing"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Pop() interface{} {
	index := len(*h) - 1
	x := (*h)[index]
	*h = (*h)[:index]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	num := x.(int)
	*h = append(*h, num)
}
func (h *IntHeap) Poll() interface{} {
	x := (*h)[0]
	*h = (*h)[1:]
	return x
}
func CutGold2(split []int) int {
	iConsume := 0
	v := IntHeap(split)
	heap.Init(&v)
	for v.Len() > 1 {
		x := v.Poll().(int)
		y := v.Poll().(int)
		iTmp := x + y
		iConsume += iTmp
		heap.Push(&v, iTmp)
	}
	return iConsume
}
func TestT2(t *testing.T) {
	arr := []int{10, 20, 30}
	gold := CutGold2(arr)
	fmt.Println(gold)
}
