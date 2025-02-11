package main

import (
	"fmt"
)

func main() {
	// Array
	var a [5]int
	a[2] = 7
	fmt.Println(a[2])
	fmt.Println(a[3])

	// Array literal
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b[2])
	fmt.Println(b[3])

	// Array length
	fmt.Println(len(b))

	// Array of arrays
	c := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(c[1][2])
	fmt.Println(c[0][1])
	fmt.Println(c[1][0])
}
