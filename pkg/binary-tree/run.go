package t_binary_tree

import "fmt"

func Run(traversalType string) {
	binaryTree := NewBinaryTree()
	binaryTree.Insert(50)
	binaryTree.Insert(25)
	binaryTree.Insert(49)
	binaryTree.Insert(75)
	binaryTree.Insert(59)
	binaryTree.Insert(85)
	binaryTree.Insert(12)
	binaryTree.Insert(28)
	binaryTree.Insert(43)
	binaryTree.Insert(30)

	fmt.Print("\nBinary tree representation: \n\n")
	binaryTree.Print("", binaryTree.root, false, true)
	fmt.Println()

	fmt.Print("Reversed binary tree representation: \n\n")
	reversedBinaryTree := binaryTree
	reversedBinaryTree.Reverse(reversedBinaryTree.root)
	reversedBinaryTree.Print("", reversedBinaryTree.root, false, true)
	fmt.Println()

	if traversalType == "preorder" {
		fmt.Println("Preorder traversal: ")
		binaryTree.PreOrderTraversal(binaryTree.root)
		return
	}
	if traversalType == "inorder" {
		fmt.Println("Inorder traversal: ")
		binaryTree.InOrderTraversal(binaryTree.root)
		return
	}
	if traversalType == "postorder" {
		fmt.Println("Postorder traversal: ")
		binaryTree.PostOrderTraversal(binaryTree.root)
		return
	}

	fmt.Println("Invalid traversal type provided - ", traversalType)
	fmt.Println("Valid traversal types are: preorder, inorder, postorder")
}
