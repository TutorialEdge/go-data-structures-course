package main

import (
	"errors"
	"fmt"
)

// Stack - our stack representation
type Stack struct {
	Elements []int
}

// Push - the method which handles pushing new elements on top
// of our stack
func (s *Stack) Push(elem int) {
	s.Elements = append(s.Elements, elem)
}

// Pop - removes the top element from the stack or returns
// an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		lastElemIndex := len(s.Elements) - 1
		var lastElement int
		lastElement, s.Elements = s.Elements[lastElemIndex], s.Elements[:lastElemIndex]
		return lastElement, nil
	}
}

// Peek - returns the top element from the stack without updating the stack
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	} else {
		return s.Elements[len(s.Elements)-1], nil
	}
}

// Length - will return the size of the stack
func (s *Stack) Length() int {
	return len(s.Elements)
}

// IsEmpty - returns true if length of stack is 0
func (s *Stack) IsEmpty() bool {
	return len(s.Elements) == 0
}

func main() {
	fmt.Println("Stacks Section")
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	peek1, _ := stack.Peek()
	fmt.Println(peek1)

	fmt.Println(stack.Length())

	elem1, _ := stack.Pop()
	fmt.Println(elem1)
	elem2, _ := stack.Pop()
	fmt.Println(elem2)
	elem3, _ := stack.Pop()
	fmt.Println(elem3)

	fmt.Println(stack.IsEmpty())
}
