package t_stack

import "fmt"

func Run() {
	fmt.Print("\nHow many elements do you want to push to the stack?  ")
	var n int
	fmt.Scan(&n)

	stack := NewStack(n)

	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Print("Enter value:  ")
		var value int
		fmt.Scan(&value)
		stack.Push(value)
	}

	fmt.Print("\nStack elements:\n\n")
	stack.Print()

	fmt.Println("Stacks are LIFO (Last In First Out) data structures, so the last element pushed to the stack is on the top.")
}
