package main

import (
	"fmt"
)

func main() {
	fmt.Println("functions in golang")
	greeter()

	// anonymous function
	func() {
		fmt.Println("anonymous function")
	}() // calling the anonymous function

	result := add(2, 3)
	fmt.Println("result is", result)

	proResult := proAdder(2, 3, 4, 5, 6)
	fmt.Println("proResult is", proResult)
}

func add(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func greeter() {
	fmt.Println("hello, world!")
}
