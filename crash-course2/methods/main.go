package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println("Structs in Go")
	// no inheritance in golang
	// no classes in golang
	// no constructors in golang
	// no generics in golang

	// creating a struct
	p1 := Person{"John", 30}
	fmt.Println(p1)
	fmt.Printf("p1 details: %+v\n", p1)

	p1.details()
}

func (p Person) details() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}
