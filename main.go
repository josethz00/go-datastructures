package main

import (
	"bufio"
	"fmt"
	t_binary_tree "go-datastructures/pkg/binary-tree"
	"os"
)

func main() {
	fmt.Println("Select a data structure to run:")
	fmt.Println("    1. Binary Tree")
	fmt.Println("    2. Linked List")
	fmt.Println("    3. Stack")
	fmt.Println("    4. Queue")
	fmt.Println("    5. Hash Table")
	fmt.Println("    6. Graph")
	fmt.Println("    7. Heap")
	fmt.Println("    8. Trie")
	fmt.Println("    9. Binary Search Tree")
	fmt.Println("    10. AVL Tree")
	fmt.Println("    11. Red-Black Tree")
	fmt.Println("    12. Map")
	fmt.Println("    13. Array")
	fmt.Println("    14. Matrix")
	fmt.Println("    15. Dynamic Array")

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No data structure selected")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	switch args[0] {
	case "1":
		fmt.Println("\nYou selected Binary Tree")
		fmt.Print("Select a traversal type:\n\n")
		fmt.Println("    - preorder")
		fmt.Println("    - inorder")
		fmt.Println("    - postorder")
		scanner.Scan()
		t_binary_tree.RunBinaryTree(scanner.Text())
	default:
		fmt.Println("TODO: Implement the rest of the data structures")
	}
}
