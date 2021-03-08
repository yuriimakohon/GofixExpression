package main

import (
	"fmt"
	"github.com/yuriimakohon/GofixExpression/postfix"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0]+` "[ infix expression ]"`)
		fmt.Println("\tExample: ", os.Args[0]+` "200 / ( 5 * 2 )"`)
		return
	}

	infixStr := strings.Join(os.Args[1:], " ")
	postfixStr, err := postfix.InfixToPostfix(infixStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(postfix.Calculate(postfixStr))
}
