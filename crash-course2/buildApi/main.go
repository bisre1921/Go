package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	return c.Title == ""
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

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, course := range courses {
		if course.ID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with the given id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if course.IsEmpty() {
		http.Error(w, "Please send a valid course", http.StatusBadRequest)
		return
	}

	// generate a random id
	course.ID = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.ID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var updatedCourse Course
			err := json.NewDecoder(r.Body).Decode(&updatedCourse)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if updatedCourse.IsEmpty() {
				http.Error(w, "Please send a valid course", http.StatusBadRequest)
				return
			}
			updatedCourse.ID = course.ID
			courses = append(courses, updatedCourse)
			json.NewEncoder(w).Encode(updatedCourse)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with the given id")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if course.ID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with the given id")
	return
}
