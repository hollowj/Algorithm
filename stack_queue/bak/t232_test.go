package bak

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

type MyQueue struct {
	stack *linkedliststack.Stack
}

func Constructor() MyQueue {
	return MyQueue{
		stack: linkedliststack.New(),
	}
}

func (this *MyQueue) Push(x int) {
	size := this.stack.Size()
	l := make([]int, size+1)
	l[size] = x
	for i := 0; i < size; i++ {
		value, _ := this.stack.Pop()
		num := value.(int)
		l[i] = num
	}

	for i := len(l) - 1; i >= 0; i-- {
		this.stack.Push(l[i])
	}
}

func (this *MyQueue) Pop() int {
	value, ok := this.stack.Pop()
	if ok {
		num := value.(int)
		return num
	}
	return 0
}

func (this *MyQueue) Peek() int {
	value, ok := this.stack.Peek()
	if ok {
		num := value.(int)
		return num
	}
	return 0
}

func (this *MyQueue) Empty() bool {
	return this.stack.Empty()
}

func TestT232(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	assert.Equal(t, 1, queue.Pop())
	queue.Push(5)
	assert.Equal(t, 2, queue.Pop())
	assert.Equal(t, 3, queue.Pop())
	assert.Equal(t, 4, queue.Pop())
	assert.Equal(t, 5, queue.Pop())
}
