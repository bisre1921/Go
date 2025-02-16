package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	ID     string `json:"string"`
	Title  string `json:"title"`
	Author Author `json:"author"`
	Price  string `json:"price"`
}

type Author struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.ID == "" && c.Title == ""
}

func main() {
	fmt.Println("courses api")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to courses api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
