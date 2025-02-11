package main

import "fmt"

func main() {
	fmt.Println("If Else")

	// if else
	if 1 > 2 {
		fmt.Println("1 > 2")
	} else {
		fmt.Println("1 <= 2")
	}

	// if else if
	if 1 > 2 {
		fmt.Println("1 > 2")
	}
	if 1 == 2 {
		fmt.Println("1 == 2")
	} else {
		fmt.Println("1 != 2")
	}

	// if else if else
	if 1 > 2 {
		fmt.Println("1 > 2")
	}
	if 1 == 2 {
		fmt.Println("1 == 2")
	} else {
		fmt.Println("1 != 2")
	}
	if 1 < 2 {
		fmt.Println("1 < 2")
	}

	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular User"
	}
	if loginCount > 10 {
		result = "Watch out"
	}
	if loginCount == 10 {
		result = "Exactly 10"
	}
	fmt.Println(result)

	// if else with initialization
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

}
