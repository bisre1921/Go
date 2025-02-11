package main

import (
	"fmt"
)

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// infinite loop
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for loop with range
	for i, day := range days {
		fmt.Println(i, day)
	}

	// break and continue
	for i, day := range days {
		if i == 2 {
			break
		}
		fmt.Println(i, day)
	}

	for i, day := range days {
		if i == 2 {
			continue
		}
		fmt.Println(i, day)
	}

	// goto
	goto mylabel
	fmt.Println("This will not be printed")
mylabel:
	fmt.Println("This will be printed")

}
