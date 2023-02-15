package expressions

import (
	"fmt"
	t_stack "go-datastructures/pkg/stack"
	number_util "go-datastructures/utils"
)

func Run() {
	fmt.Print("Select a expression type (infix/postfix):   ")
	var expressionType string
	fmt.Scan(&expressionType)

	fmt.Print("Enter your expression:   ")
	var expressionInput string
	fmt.Scan(&expressionInput)

	exprAuxStack := t_stack.NewStackStr(256)
	expression := make([]rune, len(expressionInput)*2)

	var j int = 0

	for _, curr := range expressionInput {
		fmt.Printf("%c", curr)
		if number_util.IsNumber(string(curr)) {
			exprAuxStack.Push(string(curr))
			expression[j] = curr
			j++
		} else if string(curr) == "+" || string(curr) == "-" || string(curr) == "*" || string(curr) == "/" {
			exprAuxStack.Push(string(curr))
		} else if string(curr) == ")" {
			bbvv := exprAuxStack.Pop()
			expression[j] = rune(bbvv[0])
			j++
		}
	}

	fmt.Println()
	fmt.Println("Expression: ", string(expression))
}
