package main

import (
	"fmt"
)

func main() {
	fmt.Println("golang maps")

	// create a map
	// map[key-type]value-type
	// key-type must be comparable
	// value-type can be any type

	languages := make(map[string]string)

	// add key-value pairs
	languages["JS"] = "javascript"
	languages["PY"] = "python"
	languages["RB"] = "ruby"

	fmt.Println(languages)
	fmt.Println(languages["JS"])
	fmt.Println(languages["PY"])
	fmt.Println(languages["RB"])

	// delete a key-value pair
	delete(languages, "RB")
	fmt.Println(languages)

	// check if a key exists
	_, ok := languages["PY"]
	if ok {
		fmt.Println("PY key exists")
	} else {
		fmt.Println("PY key does not exist")
	}

	// iterate over a map
	for key, value := range languages {
		fmt.Println(key, value)
	}

}
