package postfix

import (
	"github.com/golang-collections/collections/stack"
	"strconv"
	"strings"
)

func doMath(stack *stack.Stack, op byte) error {
	if stack.Len() < 2 {
		return ErrInvalidPostfixExpr
	}
	s, f := stack.Pop().(float64), stack.Pop().(float64)

	switch op {
	case '+':
		stack.Push(f + s)
	case '-':
		stack.Push(f - s)
	case '*':
		stack.Push(f * s)
	case '/':
		stack.Push(f / s)
	}
	return nil
}

func Calculate(s string) (float64, error) {
	tokens := strings.Split(s, " ")
	operandStack := stack.New()

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "/" || token == "*" {
			if err := doMath(operandStack, token[0]); err != nil {
				return 0, err
			}
		} else {
			operand, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			operandStack.Push(operand)
		}
	}
	return operandStack.Pop().(float64), nil
}
