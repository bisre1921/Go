package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, Go Time!")
	presentTime := time.Now()
	fmt.Println("The time is", presentTime)
	fmt.Println(presentTime.Format("2006-01-02 15:04:05"))

	createdDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created date:", createdDate.Format("2006-01-02 15:04:05"))
}
