package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza shop")
	fmt.Println("please rate our pizza")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	fmt.Println("Thanks for rating our pizza", input)

	// convert the input to float
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("An error occured while converting the input to float. Please try again", err)
		return
	} else {
		fmt.Println("Added 1 to your rating", numRating+1)
	}

}
