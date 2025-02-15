package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string `json:"website"`
	Password string
	Tags     []string
}

func main() {
	fmt.Println("json")
	EncodeJson()
}

func EncodeJson() {
	c := []course{
		{"Docker", 100, "Udemy", "password", []string{"docker", "container"}},
		{"Kubernetes", 100, "Udemy", "password", []string{"k8s", "container"}},
		{"Jenkins", 100, "Udemy", "password", []string{"ci", "cd"}},
	}

	// package this data into json
	finalJson, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", finalJson)
}
