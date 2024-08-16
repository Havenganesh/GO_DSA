package main

type Node1 struct {
	Value int
	Left  *Node1
	Right *Node1
}

func (n *Node1) Insert(Value int) {
	if Value <= n.Value {
		if n.Left == nil {
			n.Left = &Node1{}
		}
		n.Left.Value = Value
	} else {
		if n.Right == nil {
			n.Right = &Node1{}
		}
		n.Right.Value = Value
	}
}

func (n *Node1) Search(value int) bool {
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
