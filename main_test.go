package main_test

import (
	"fmt"
	"testing"

	main "github.com/lfernandezlo/go/binary-tree"
)

func TestBinaryTreePrint(t *testing.T) {
	tree := main.CreateTree(main.Node{
		Data: 10,
		LeftNode: &main.Node{
			Data: 5,
			LeftNode: &main.Node{
				Data: 2,
			},
			RightNode: &main.Node{
				Data: 8,
			},
		},
		RightNode: &main.Node{
			Data: 15,
			LeftNode: &main.Node{
				Data: 11,
			},
			RightNode: &main.Node{
				Data: 18,
			},
		},
	})

	out := tree.Node.Print("")

	if out != "2,5,8,10,11,15,18," {
		t.Error(fmt.Sprintf("Should be equal to 2,5,8,10,11,15,18,. Got %v", out))
	}
}

func TestBinaryTreeInsertLeft(t *testing.T) {
	tree := main.CreateTree(main.Node{
		Data: 10,
		LeftNode: &main.Node{
			Data: 5,
		},
		RightNode: &main.Node{
			Data: 15,
		},
	})

	tree.Node.Insert(8)

	if tree.Node.LeftNode.RightNode == nil || tree.Node.LeftNode.RightNode.Data != 8 {
		t.Error("Node has not been inserted")
	}
}

func TestBinaryTreeInsertRight(t *testing.T) {
	tree := main.CreateTree(main.Node{
		Data: 10,
		LeftNode: &main.Node{
			Data: 5,
		},
		RightNode: &main.Node{
			Data: 15,
		},
	})

	tree.Node.Insert(21)

	if tree.Node.RightNode.RightNode == nil || tree.Node.RightNode.RightNode.Data != 21 {
		t.Error("Node has not been inserted")
	}
}

func TestBinaryTreeInsertNoLeftNode(t *testing.T) {
	tree := main.CreateTree(main.Node{
		Data: 10,
		RightNode: &main.Node{
			Data: 15,
		},
	})

	tree.Node.Insert(5)

	if tree.Node.LeftNode == nil || tree.Node.LeftNode.Data != 5 {
		t.Error("Node has not been inserted")
	}
}

func TestBinaryTreeInsertNoRightNode(t *testing.T) {
	tree := main.CreateTree(main.Node{
		Data: 10,
	})

	tree.Node.Insert(15)

	if tree.Node.RightNode == nil || tree.Node.RightNode.Data != 15 {
		t.Error("Node has not been inserted")
	}
}

func TestBinaryTreeContains(t *testing.T) {
	tree := main.CreateTree(main.Node{
		Data: 10,
		LeftNode: &main.Node{
			Data: 5,
		},
		RightNode: &main.Node{
			Data: 15,
		},
	})

	contains := tree.Node.Contains(5)

	if !contains {
		t.Error("It should contain node 5")
	}
}

func TestBinaryTreeDoesNotContainsLeft(t *testing.T) {
	tree := main.CreateTree(main.Node{
		Data: 10,
		RightNode: &main.Node{
			Data: 15,
		},
	})

	contains := tree.Node.Contains(5)

	if contains {
		t.Error("It should not contain node 5")
	}
}

func TestBinaryTreeDoesNotContainsRight(t *testing.T) {
	tree := main.CreateTree(main.Node{
		Data: 10,
		LeftNode: &main.Node{
			Data: 5,
		},
	})

	contains := tree.Node.Contains(15)

	if contains {
		t.Error("It should not contain node 15")
	}
}

func TestBinaryTreeContainsRightRecursively(t *testing.T) {
	tree := main.CreateTree(main.Node{
		Data: 10,
		RightNode: &main.Node{
			Data: 15,
			RightNode: &main.Node{
				Data: 21,
			},
		},
	})

	contains := tree.Node.Contains(21)

	if !contains {
		t.Error("It should contain 21")
	}
}
