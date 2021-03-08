package postfix

import (
	"errors"
	"fmt"
)

var (
	ErrParenthesesSeq     = errors.New("Invalid parentheses placing\n")
	ErrInvalidPostfixExpr = errors.New("Invalid postfix expression\n")
)

type ErrToken struct {
	message string
}

func (e ErrToken) Error() string {
	return fmt.Sprintf("Token: %s\n", e.message)
}

func NewErrToken(message string) error {
	return &ErrToken{
		message: message,
	}
}
