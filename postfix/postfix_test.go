package postfix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	infix := []string{
		"A + B",
		"AaA - bBb * CcC",
		"1 / ( 2 + 3 )",
		"5 + 185 * 756 / ( 125 - 11 )",
		"( ( 5 + 185 ) * 756 ) / ( 125 - 11 ) + 4 - ( 21 / 7 )",
	}
	expects := []string{
		"A B +",
		"AaA bBb CcC * -",
		"1 2 3 + /",
		"5 185 756 * 125 11 - / +",
		"5 185 + 756 * 125 11 - / 4 + 21 7 / -",
	}

	for i, expect := range expects {
		actual, err := InfixToPostfix(infix[i])
		assert.NoError(t, err)
		assert.Equal(t, expect, actual)
	}

	_, err := InfixToPostfix("( A + B (")
	assert.ErrorIs(t, ErrParenthesesSeq, err)

	_, err = InfixToPostfix("A + B )")
	assert.ErrorIs(t, ErrParenthesesSeq, err)
}

func TestCalculate(t *testing.T) {
	input := []string{
		"55 5 +",
		"100 3 /",
		"1 97 3 + /",
		"5 185 756 * 125 11 - / +",
		"5 185 + 756 * 125 11 - / 4 + 21 7.5 / -",
	}
	var expects = []float64{
		55 + 5,
		float64(100) / float64(3),
		float64(1) / float64(100),
		float64(5) + float64(185)*float64(756)/float64(125-11),
		float64((5+185)*756)/float64(125-11) + float64(4) - (float64(21) / (7.5)),
	}
	for i, expect := range expects {
		actual, err := Calculate(input[i])
		assert.NoError(t, err)
		assert.Equal(t, expect, actual)
	}
}
