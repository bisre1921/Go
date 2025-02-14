package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("files in golang")

	content := "Hello, World!"

	file, err := os.Create("./hello.txt")
	if err != nil {
		fmt.Println("Error creating file")
		return
	}
	defer file.Close()

	len, err := file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file")
		return
	}
	fmt.Println(len, "bytes written successfully")

	readFile("./hello.txt")

}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	fmt.Println("Data read from file:", string(dataByte))

}
