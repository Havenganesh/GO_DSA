package main

import (
	"fmt"
)

// Node represents a node in the BST
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert adds a new node to the BST
func (n *Node) Insert(value int) {
	if value <= n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search looks for a node with a specific value
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}
	if n.Value == value {
		return true
	}
	if value < n.Value {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

// InOrderTraversal traverses the tree in order
func (n *Node) InOrderTraversal() {
	if n != nil {
		n.Left.InOrderTraversal()
		fmt.Printf("%d ", n.Value)
		n.Right.InOrderTraversal()
	}
}

func main() {
	root := &Node{Value: 10}
	root.Insert(5)
	root.Insert(15)
	root.Insert(3)
	root.Insert(7)
	root.Insert(13)
	root.Insert(17)

	fmt.Print("InOrder Traversal: ")
	root.InOrderTraversal() // Output: 3 5 7 10 13 15 17
	fmt.Println()

	fmt.Println("Search 7:", root.Search(7))   // Output: true
	fmt.Println("Search 12:", root.Search(12)) // Output: false
}
