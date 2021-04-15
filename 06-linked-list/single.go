package main

import "fmt"

// LinkedList
type LinkedList struct {
	Head *Node
	Size int
}

// Node
type Node struct {
	Next  *Node
	Value string
}

func (l *LinkedList) Insert(elem string) {
	node := Node{
		Next:  l.Head,
		Value: elem,
	}
	l.Head = &node
	l.Size++
}

func (l *LinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

func (l *LinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Println("%+v\n", current)
		current = current.Next
	}
}

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
