package stack_queue

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getOpNum(stack []string) (int, int) {
	size := len(stack)
	num1Str := stack[size-2]
	num2Str := stack[size-1]
	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(num2Str)
	if err != nil {
		panic(err)
	}
	return num1, num2
}

//	func evalRPN(tokens []string) int {
//		stack := make([]string, 0)
//		for i := 0; i < len(tokens); i++ {
//			var token = tokens[i]
//			size := len(stack)
//			if token == "+" {
//				num1, num2 := getOpNum(stack)
//				r := num1 + num2
//				token = strconv.Itoa(r)
//				stack = stack[:size-2]
//			} else if token == "-" {
//				num1, num2 := getOpNum(stack)
//				r := num1 - num2
//				token = strconv.Itoa(r)
//				stack = stack[:size-2]
//			} else if token == "*" {
//				num1, num2 := getOpNum(stack)
//				r := num1 * num2
//				token = strconv.Itoa(r)
//				stack = stack[:size-2]
//			} else if token == "/" {
//				num1, num2 := getOpNum(stack)
//				r := num1 / num2
//				token = strconv.Itoa(r)
//				stack = stack[:size-2]
//			}
//			stack = append(stack, token)
//		}
//		num, err := strconv.Atoi(stack[0])
//		if err != nil {
//			panic(err)
//		}
//		return num
//	}
func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		var token = tokens[i]
		num, err := strconv.Atoi(token)
		if err != nil {
			size := len(stack)
			num1 := stack[size-2]
			num2 := stack[size-1]
			stack = stack[:size-2]
			if token == "+" {
				stack = append(stack, num1+num2)
			} else if token == "-" {
				stack = append(stack, num1-num2)
			} else if token == "*" {
				stack = append(stack, num1*num2)
			} else if token == "/" {
				stack = append(stack, num1/num2)
			}
		} else {
			stack = append(stack, num)
		}
	}
	return stack[0]
}

func TestT150(t *testing.T) {
	assert.Equal(t, 9, evalRPN([]string{"2", "1", "+", "3", "*"}))
	assert.Equal(t, 6, evalRPN([]string{"4", "13", "5", "/", "+"}))
	assert.Equal(t, 22, evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
