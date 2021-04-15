package main

import "fmt"

// LinkedList - this has to store both the head of our linked list
// as well as the size. Whenever we perform an action such as insert or
// delete, we need to take care to manually update this Size field.
type LinkedList struct {
	Head *Node
	Size int
}

// Node - the representation of a Node within our linked list object
type Node struct {
	Next  *Node
	Value string
}

// Insert - the insert method takes in an element and
// creates a new Node using this value. It then sets this node as the
// first element in our linked list.
//
// This is an implementation of a linked list which follows the LIFO
// access pattern. The last element added becomes the head, and if we
// read our Linked list in order, the head will be the first element
// we read.
func (l *LinkedList) Insert(elem string) {
	node := Node{
		Next:  l.Head,
		Value: elem,
	}
	l.Head = &node
	l.Size++
}

// DeleteFirst - deletes the first element from the head of our
// linked list and also updates the size. By simply removing the
// reference to it, the internal Go garbage collector will clean
// this up for us!
func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

// List - iterates through every element within our linked list
// using the same LIFO access pattern we talked about before.
func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Println("%+v\n", current)
		current = current.Next
	}
}

// Search - iterates through every node within the linked list until
// it reaches the end of the linked list. If any of the Nodes contain the
// inputted string then it will return the pointer to that particular node.
func (l *LinkedList) Search(elem string) *Node {
	current := l.Head
	for current != nil {
		if current.Value == elem {
			return current
		}
		current = current.Next
	}
	return nil
}

// Delete - performs the action of removing a node from our
// linked list if it matches the inputted string.
//
// This needs to keep track of the previous element and the current
// element so that it can effectively relink the previous element
// to the next element after the node that is to be deleted.
//
// e.g. Delete("C")
// ["A" => "B" => "C" => "D"]
// In this example when we iterate through the linked list we find "C"
// we have 'previous' = "B" and 'current' = "C". We then perform the action of
// linking 'previous' = "B" to current.Next which is equal to "D".
// this leaves us with ["A" => "B" => "D"]
func (l *LinkedList) Delete(elem string) {
	previous := l.Head
	current := l.Head
	for current != nil {
		if current.Value == elem {
			previous.Next = current.Next
		}
		previous = current
		current = current.Next
	}
}
