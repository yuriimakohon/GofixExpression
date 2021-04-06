package postfix

import (
	"github.com/golang-collections/collections/stack"
	"regexp"
	"strings"
)

var (
	prec = map[byte]uint8{
		'*': 3,
		'/': 3,
		'+': 2,
		'-': 2,
		'(': 1,
	}

	operatorRegexp = regexp.MustCompile(`^[()/*+-]$`)
	operandRegexp  = regexp.MustCompile(`^[A-Za-z0-9]*$`)
)

func pushOperator(stack *stack.Stack, operator byte, resList *[]string) error {
	switch operator {
	case '(':
		stack.Push(operator)
	case ')':
		{
			for curr := stack.Pop(); ; curr = stack.Pop() {
				if curr == nil {
					return ErrParenthesesSeq
				}
				if curr.(byte) == '(' {
					break
				}
				*resList = append(*resList, string(curr.(byte)))
			}
		}
	default:
		{
			for stack.Len() > 0 && prec[stack.Peek().(byte)] >= prec[operator] {
				*resList = append(*resList, string(stack.Pop().(byte)))
			}
			stack.Push(operator)
		}
	}
	return nil
}

func InfixToPostfix(infix string) (string, error) {
	tokens := strings.Split(infix, " ")
	operatorStack := stack.New()

	var resList []string

	for _, token := range tokens {
		if operandRegexp.MatchString(token) {
			resList = append(resList, token)
		} else if operatorRegexp.MatchString(token) {
			if err := pushOperator(operatorStack, token[0], &resList); err != nil {
				return "", err
			}
		} else {
			return "", NewErrToken("invalid token format: " + token)
		}
	}

	for curr := operatorStack.Pop(); curr != nil; curr = operatorStack.Pop() {
		if curr.(byte) == '(' || curr.(byte) == ')' {
			return "", ErrParenthesesSeq
		}
		resList = append(resList, string(curr.(byte)))
	}

	return strings.Join(resList, " "), nil
}
