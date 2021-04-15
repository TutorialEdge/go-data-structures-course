package main

import (
	"errors"
	"fmt"
)

// Set - our representation of a set data structure
type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

// Add - adds an element to our Set
func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

// Delete - removes an element from our set if it exists
func (s *Set) Delete(elem string) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("element not present in set")
	}
	delete(s.Elements, elem)
	return nil
}

// Contains - checks to see if an element exists within the set
func (s *Set) Contains(elem string) bool {
	_, exists := s.Elements[elem]
	return exists
}

// List - prints out all of the elements within our set
func (s *Set) List() {
	for k, _ := range s.Elements {
		fmt.Println(k)
	}
}

func main() {
	fmt.Println("Sets Tutorial")
	mySet := NewSet()

	mySet.Add("Earth")
	mySet.Add("Venus")
	mySet.Add("Mars")
	mySet.Add("Earth")
	mySet.Delete("Venus")
	mySet.List()

	fmt.Println(mySet.Contains("Mars"))
}
