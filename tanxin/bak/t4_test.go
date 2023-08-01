package bak

import (
	"fmt"
	"testing"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

type MedianFinder struct {
	minQueue *binaryheap.Heap
	maxQueue *binaryheap.Heap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minQueue: binaryheap.NewWith(func(a, b interface{}) int {
			return utils.IntComparator(a, b)
		}),
		maxQueue: binaryheap.NewWith(func(a, b interface{}) int {
			return -utils.IntComparator(a, b)
		}),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxQueue.Size() > 0 {
		value, _ := this.maxQueue.Peek()
		if value.(int) < num {
			this.minQueue.Push(num)
		} else {
			this.maxQueue.Push(num)
		}
	} else {
		this.maxQueue.Push(num)
	}
	if this.maxQueue.Size() >= this.minQueue.Size()+2 {
		value, _ := this.maxQueue.Pop()
		this.minQueue.Push(value)
	}
	if this.maxQueue.Size()+2 <= this.minQueue.Size() {
		value, _ := this.minQueue.Pop()
		this.maxQueue.Push(value)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minQueue.Size() == this.maxQueue.Size() {
		var x int
		var y int
		a, ok := this.minQueue.Peek()
		if ok {
			x = a.(int)
		}
		b, ok := this.maxQueue.Peek()
		if ok {
			y = b.(int)
		}
		return float64((x + y)) / 2
	} else {
		if this.minQueue.Size() > this.maxQueue.Size() {
			value, ok := this.minQueue.Peek()
			if ok {
				return float64(value.(int))
			}
		} else {
			value, ok := this.maxQueue.Peek()
			if ok {
				return float64(value.(int))
			}
		}
	}

	return 0
}
func TestT4(t *testing.T) {
	constructor := Constructor()
	constructor.AddNum(1)
	fmt.Println(constructor.FindMedian())
	constructor.AddNum(2)
	fmt.Println(constructor.FindMedian())
	constructor.AddNum(3)
	fmt.Println(constructor.FindMedian())
	constructor.AddNum(4)
	fmt.Println(constructor.FindMedian())
	constructor.AddNum(5)
	fmt.Println(constructor.FindMedian())
	constructor.AddNum(6)
	fmt.Println(constructor.FindMedian())
	constructor.AddNum(7)
	fmt.Println(constructor.FindMedian())
	constructor.AddNum(8)
	fmt.Println(constructor.FindMedian())
	constructor.AddNum(9)
	fmt.Println(constructor.FindMedian())
	constructor.AddNum(10)
	fmt.Println(constructor.FindMedian())
}
