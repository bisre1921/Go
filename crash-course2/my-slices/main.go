package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices")

	var fruits = []string{"apple", "banana", "grape", "orange"}
	fmt.Printf("fruits: %T\n", fruits)
	fruits = append(fruits, "mango")
	fmt.Println(fruits)

	// Slicing
	fmt.Println(fruits[1:3])
	fmt.Println(fruits[:2])
	fmt.Println(fruits[2:])
	fmt.Println(fruits[:])

	// Length and Capacity
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	// Make
	numbers := make([]int, 5, 10)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

}
