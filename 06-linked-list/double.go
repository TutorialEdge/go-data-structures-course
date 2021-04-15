package main

import "fmt"

// DoublyLinkedList
type DoublyLinkedList struct {
	Head *DoubleNode
	Tail *DoubleNode
	Size int
}

// DoubleNode
type DoubleNode struct {
	Next  *DoubleNode
	Prev  *DoubleNode
	Value string
}

func (l *DoublyLinkedList) Insert(elem string) {
	node := DoubleNode{
		Next:  l.Head,
		Prev:  nil,
		Value: elem,
	}
	l.Head = &node
	l.Tail = &node
	l.Size++
}

func (l *DoublyLinkedList) DeleteFirst() {
	l.Head = l.Head.Next
	l.Size--
}

func (l *DoublyLinkedList) List() {
	current := l.Head
	for current != nil {
		fmt.Println("%+v\n", current)
		current = current.Next
	}
}

func (l *DoublyLinkedList) Search(elem string) *DoubleNode {
	current := l.Head
	for current != nil {
		if current.Value == elem {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *DoublyLinkedList) Delete(elem string) {
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
