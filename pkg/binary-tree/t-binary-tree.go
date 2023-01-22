package t_binary_tree

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(value int) {
	if t.root == nil { // if the tree is empty then the root will be nil
		t.root = &Node{value: value} // in this case we create a new node and assign it to the root
	} else { // if the tree is not empty then we need to find the right place for the new node
		currentNode := t.root // we start from the root

		for {
			if value < currentNode.value { // in binary trees the left node is always smaller than the right node
				if currentNode.left == nil { // if the left node is nil then we can insert the new node here
					currentNode.left = &Node{value: value}
					return
				}
				currentNode = currentNode.left // if the left node is not nil then we need to go deeper, so we assign the left node to the current node
			} else { // if `value` is greater than the current node value then we need to go to the right
				if currentNode.right == nil { // if the right node is nil then we can insert the new node here
					currentNode.right = &Node{value: value}
					return
				}
				currentNode = currentNode.right // if the right node is not nil then we need to go deeper, so we assign the right node to the current node
			}
		}
	}
}

func (t *BinaryTree) PreOrderTraversal(currentNode *Node) {
	if currentNode == nil {
		return
	}

	fmt.Print(currentNode.value, " ")      // print the current node value
	t.PreOrderTraversal(currentNode.left)  // recursively call the function for the left node
	t.PreOrderTraversal(currentNode.right) // recursively call the function for the right node
}
