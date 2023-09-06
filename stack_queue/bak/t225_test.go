package bak

import (
	"testing"

	"github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/stretchr/testify/assert"
)

type MyStack struct {
	queueIn *linkedlistqueue.Queue
}

func Constructor1() MyStack {
	return MyStack{
		queueIn: linkedlistqueue.New(),
	}
}

func (this *MyStack) Push(x int) {
	this.queueIn.Enqueue(x)
}

func (this *MyStack) OpTop(pop bool) int {
	size := this.queueIn.Size()
	for size > 1 {
		size--
		value, _ := this.queueIn.Dequeue()
		x := value.(int)
		this.queueIn.Enqueue(x)
	}
	value, _ := this.queueIn.Dequeue()
	if !pop {
		this.queueIn.Enqueue(value)
	}
	x := value.(int)
	return x
}
func (this *MyStack) Pop() int {
	return this.OpTop(true)
}

func (this *MyStack) Top() int {
	return this.OpTop(false)

}

func (this *MyStack) Empty() bool {
	return this.queueIn.Empty()
}
func TestT225(t *testing.T) {
	stack := Constructor1()
	stack.Push(1)
	stack.Push(2)
	assert.Equal(t, 2, stack.Top())
	assert.Equal(t, 2, stack.Pop())
	assert.Equal(t, 1, stack.Pop())
	assert.Equal(t, true, stack.Empty())
}
