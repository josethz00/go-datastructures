package main

import (
	"bufio"
	"fmt"
	t_stack "go-datastructures/pkg/stack"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a string:  ")
	scanner.Scan()
	var input string = scanner.Text()

	stackA := t_stack.NewStackRune(len(input))

	for i := 0; i < len(input); i++ {
		stackA.Push(rune(input[i]))
	}

	for !stackA.IsEmpty() {
		current, err := stackA.Pop()

		if err == nil {
			fmt.Printf("%c", current)
		}
	}
}
