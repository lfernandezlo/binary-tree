package main

import (
	"fmt"
)

// Print
// Insert
// Contains

type Tree struct {
	Node Node
}

type Node struct {
	Data      int
	LeftNode  *Node
	RightNode *Node
}

func (n *Node) Print(s string) string {
	var left, center, right string

	if n.LeftNode != nil {
		left = n.LeftNode.Print(s)
	}

	center = left + fmt.Sprintf("%v,", n.Data)

	if n.RightNode != nil {
		right = n.RightNode.Print(center)
	}

	return center + right
}

func (n *Node) Insert(i int) {
	if i < n.Data {
		if n.LeftNode == nil {
			n.LeftNode = &Node{
				Data: i,
			}
		} else {
			n.LeftNode.Insert(i)
		}
	} else {
		if n.RightNode == nil {
			n.RightNode = &Node{
				Data: i,
			}
		} else {
			n.RightNode.Insert(i)
		}
	}
}

func (n *Node) Contains(i int) bool {
	if i == n.Data {
		return true
	}

	if i < n.Data {
		if n.LeftNode == nil {
			return false
		}

		return n.LeftNode.Contains(i)
	} else {
		if n.RightNode == nil {
			return false
		}

		return n.RightNode.Contains(i)
	}
}

func main() {
	t := CreateTree(Node{
		Data: 10,
		LeftNode: &Node{
			Data: 5,
		},
		RightNode: &Node{
			Data: 15,
		},
	})

	println("Previously contains 8", t.Node.Contains(8))
	println("Inserting 8")
	t.Node.Insert(8)
	println("Contains 8", t.Node.Contains(8))
	println("Tree", t.Node.Print(""))
}

func CreateTree(n Node) *Tree {
	return &Tree{Node: n}
}
