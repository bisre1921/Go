package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hello, " + input)
		fmt.Printf("Type of input is %T\n", input)
	}
}
