package main

import (
	"errors"
	"fmt"
)

// Queue - our representation of a queue data structure
type Queue struct {
	Elements []int
}

// Enqueue - add an element of type int to the end of our queue
func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
}

// Dequeue - returns the first element from our queue
func (q *Queue) Dequeue() (int, error) {
	// return the first element
	// update elements slice
	// validate queue is not empty
	if len(q.Elements) == 0 {
		return 0, errors.New("empty queue")
	}
	var firstElement int
	firstElement, q.Elements = q.Elements[0], q.Elements[1:]
	return firstElement, nil
}

// Length - returns the length of our queue
func (q *Queue) Length() int {
	return len(q.Elements)
}

// Peek - returns the first element from our queue without updating queue
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}

// IsEmpty - return true if queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Length() == 0
}

func main() {
	fmt.Println("Queues Section")

	queue := Queue{}
	fmt.Println(queue)
	queue.Enqueue(1)
	fmt.Println(queue)
	queue.Enqueue(2)
	fmt.Println(queue)
	elem, _ := queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue)

}
