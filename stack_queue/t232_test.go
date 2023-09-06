package stack_queue

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

type MyQueue struct {
	stackIn  *linkedliststack.Stack
	stackOut *linkedliststack.Stack
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  linkedliststack.New(),
		stackOut: linkedliststack.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.stackOut.Size() == 0 {
		for !this.stackIn.Empty() {
			value, _ := this.stackIn.Pop()
			num := value.(int)
			this.stackOut.Push(num)
		}
	}
	value, _ := this.stackOut.Pop()
	return value.(int)
}

func (this *MyQueue) Peek() int {
	x := this.Pop()
	this.stackOut.Push(x)
	return x
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.Empty() && this.stackOut.Empty()
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
func TestT232_1(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	assert.Equal(t, 1, queue.Peek())
	assert.Equal(t, 1, queue.Pop())

	assert.Equal(t, false, queue.Empty())
}
