package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	// Declaring a variable
	var a int = 42
	fmt.Println(a)

	// Declaring a pointer
	var b *int = &a
	fmt.Println(b)

	// Dereferencing a pointer
	fmt.Println(*b)

	// Changing the value of a variable through a pointer
	*b = 100
	fmt.Println(a)

	// Pointer arithmetic
	// Go does not support pointer arithmetic
	// b = b + 1 // This will throw an error
	// b = b - 1 // This will throw an error
	// b = b * 1 // This will throw an error
	// b = b / 1 // This will throw an error

	// Pointers to pointers
	var c int = 200
	var d *int = &c
	var e **int = &d
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(*d)
	fmt.Println(**e)

}
