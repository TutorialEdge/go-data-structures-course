package main

import "fmt"

type Node struct {
	left  *Node
	data  int
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

// Method for BinarySearchTree struct
func (tree *BinarySearchTree) insert(data int) *BinarySearchTree {
	if tree.root == nil {
		tree.root = &Node{data: data, left: nil, right: nil}
	} else {
		tree.root.insert(data)
	}
	return tree
}

// Method for Node struct
func (node *Node) insert(data int) {
	if node == nil {
		return
	} else if data < node.data {
		if node.left == nil {
			node.left = &Node{data: data, left: nil, right: nil}
		} else {
			node.left.insert(data)
		}
	} else {
		if node.right == nil {
			node.right = &Node{data: data, left: nil, right: nil}
		} else {
			node.right.insert(data)
		}
	}
}

// InOrder Traversal
// Kind of sorted in ascending order
func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Print(root.data, " ")
	inOrder(root.right)
}

// InOrder Traversal
func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.data, " ")
	inOrder(root.left)
	inOrder(root.right)
}

/*
Just for tryouts
func main() {
	tree := &BinarySearchTree{}
	tree.insert(10)
	tree.insert(5)
	tree.insert(15)
	tree.insert(4)
	tree.insert(16)
	tree.insert(6)
	tree.insert(13)
	tree.insert(1)
	tree.insert(12)
	tree.insert(2)
	tree.insert(20)
	inOrder(tree.root)
	fmt.Println()
	preOrder(tree.root)
}
*/
