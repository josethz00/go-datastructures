package t_binary_tree

import "fmt"

func RunBinaryTree(traversalType string) {
	binaryTree := NewBinaryTree()
	binaryTree.Insert(50)
	binaryTree.Insert(25)
	binaryTree.Insert(75)
	binaryTree.Insert(12)
	binaryTree.Insert(37)
	binaryTree.Insert(43)
	binaryTree.Insert(30)

	if traversalType == "preorder" {
		binaryTree.PreOrderTraversal(binaryTree.root)
		return
	}
	if traversalType == "inorder" {
		// binaryTree.InOrderTraversal(binaryTree.root)
		return
	}
	if traversalType == "postorder" {
		// binaryTree.PostOrderTraversal(binaryTree.root)
		return
	}

	fmt.Println("Invalid traversal type provided - ", traversalType)
	fmt.Println("Valid traversal types are: preorder, inorder, postorder")
}
