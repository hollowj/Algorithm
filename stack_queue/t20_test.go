package stack_queue

import (
	"testing"

	"github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/stretchr/testify/assert"
)

func isValid(s string) bool {
	m := make(map[int32]int32)
	m[')'] = '('
	m['}'] = '{'
	m[']'] = '['
	stack := linkedliststack.New()
	for _, ch := range s {
		switch ch {
		case ')', '}', ']':
			value, ok := stack.Peek()
			if ok {
				topCh := value.(int32)
				if topCh == m[ch] {
					stack.Pop()
					continue
				} else {
					return false
				}
			}
			fallthrough
		default:
			stack.Push(ch)

		}
	}
	return stack.Empty()
}

func TestT20(t *testing.T) {
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, true, isValid("()[]{}"))
	assert.Equal(t, false, isValid("(]"))
}
