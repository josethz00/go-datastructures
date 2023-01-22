package t_binary_tree

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
			if value < currentNode.value {
				if currentNode.left == nil {
					currentNode.left = &Node{value: value}
					return
				}
				currentNode = currentNode.left
			} else {
				if currentNode.right == nil {
					currentNode.right = &Node{value: value}
					return
				}
				currentNode = currentNode.right
			}
		}
	}
}
