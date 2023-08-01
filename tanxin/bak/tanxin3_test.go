package bak

import (
	"container/heap"
	"fmt"
	"testing"
)

type Project struct {
	cost int
	gain int
}
type CostHeap []Project

func NewCostHeap(projects []Project) *CostHeap {
	h := CostHeap(projects)
	hP := &h
	heap.Init(hP)
	return hP
}
func (h CostHeap) Len() int {
	return len(h)
}
func (h CostHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}
func (h CostHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *CostHeap) Pop() interface{} {
	index := len(*h) - 1
	x := (*h)[index]
	*h = (*h)[:index]
	return x
}
func (h *CostHeap) Push(x interface{}) {
	num := x.(Project)
	*h = append(*h, num)
}
func (h *CostHeap) Poll() interface{} {
	x := (*h)[0]
	*h = (*h)[1:]
	return x
}
func (h *CostHeap) First() Project {
	x := (*h)[0]
	return x
}

type GainHeap []Project

func NewGainHeap(projects []Project) *GainHeap {
	h := GainHeap(projects)
	hP := &h
	heap.Init(hP)
	return hP
}
func (h GainHeap) Len() int {
	return len(h)
}
func (h GainHeap) Less(i, j int) bool {
	return h[i].gain < h[j].gain
}
func (h GainHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *GainHeap) Pop() interface{} {
	index := len(*h) - 1
	x := (*h)[index]
	*h = (*h)[:index]
	return x
}
func (h *GainHeap) Push(x interface{}) {
	num := x.(Project)
	*h = append(*h, num)
}
func (h *GainHeap) Poll() interface{} {
	x := (*h)[0]
	*h = (*h)[1:]
	return x
}
func (h *GainHeap) Peek() Project {
	x := (*h)[0]
	return x
}
func plan(projects []Project, own int, k int) int {
	costHeap := NewCostHeap(projects)
	gainHeap := NewGainHeap([]Project{})

	for k > 0 {
		for costHeap.Len() > 0 {
			if costHeap.First().cost <= own {
				gainHeap.Push(costHeap.Poll())
			} else {
				break
			}
		}
		if gainHeap.Len() == 0 {
			return own
		}
		own += gainHeap.Pop().(Project).gain
		k -= 1
	}
	return own
}
func TestLiRun(t *testing.T) {
	projects := []Project{
		{cost: 1, gain: 1},
		{cost: 1, gain: 4},
		{cost: 2, gain: 3},
		{cost: 2, gain: 7},
		{cost: 3, gain: 2},
		{cost: 4, gain: 10},
	}
	own := 1
	k := 4
	i := plan(projects, own, k)
	fmt.Println(i)
}
