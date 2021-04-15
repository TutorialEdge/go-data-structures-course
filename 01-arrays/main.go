package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	myArray := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(myArray[7])
	for i, n := range myArray {
		fmt.Printf("Index: %d\n", i)
		fmt.Println(n)
	}
}
