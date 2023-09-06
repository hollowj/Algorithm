package stack_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//func removeDuplicates(s string) string {
//	stack := linkedliststack.New()
//	for _, ch := range s {
//		value, ok := stack.Peek()
//		if ok {
//			x := value.(int32)
//			if x == ch {
//				stack.Pop()
//				continue
//			}
//		}
//		stack.Push(ch)
//	}
//	int32s := make([]int32, stack.Size())
//	for i := stack.Size() - 1; i >= 0; i-- {
//		value, _ := stack.Pop()
//		int32s[i] = value.(int32)
//	}
//	return string(int32s)
//}

func removeDuplicates(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
func TestT1047(t *testing.T) {
	assert.Equal(t, "ca", removeDuplicates("abbaca"))
}
