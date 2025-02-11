package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// switch case
	// switch case is a control structure that allows you to test multiple conditions
	// and execute a block of code based on the first condition that is true
	// switch case is a good alternative to if-else if-else chain

	// syntax
	// switch expression {
	// case value1:
	//		// code block
	// case value2:
	//		// code block
	// default:
	//		// code block
	// }

	// example
	// Generate a random number between 1 and 6
	randomNumber := rand.Intn(6) + 1
	fmt.Println("Random number:", randomNumber)

	switch randomNumber {
	case 1:
		fmt.Println("dice value is 1, you can move 1 step")
	case 2:
		fmt.Println("dice value is 2, you can move 2 steps")
	case 3:
		fmt.Println("dice value is 3, you can move 3 steps")
	case 4:
		fmt.Println("dice value is 4, you can move 4 steps")
		fallthrough
	case 5:
		fmt.Println("dice value is 5, you can move 5 steps")
	case 6:
		fmt.Println("dice value is 6, you can move 6 steps")
	default:
		fmt.Println("invalid dice value")

	}
}
