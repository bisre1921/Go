package main

import (
	"log"      // Provides logging utilities for debugging and error handling. Used to log messages, warnings, and errors
	"net/http" // Provides functionality for building HTTP servers and making HTTP requests. Used to create web servers, handle requests, and communicate over HTTP
)

func main() {
	fileServer := http.FileServer(http.Dir("./static")) // Create a file server that serves files from the static directory
	http.Handle("/", fileServer)                        // Handle all requests to the root URL with the file server
	http.HandleFunc("/hello", helloHandler)             // Handle all requests to the /hello URL with the helloHandler function
	http.HandleFunc("/form", formHandler)               // Handle all requests to the /form URL with the formHandler function

	log.Println("Starting server on :8080...") // Log a message to indicate that the server is starting
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	w.Write([]byte("Hello, world!")) // Write a response to the client
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Cannot parse form.", http.StatusNotFound)
		return
	}
	w.Write([]byte("form submitted"))
	name := r.FormValue("name")
	address := r.FormValue("address")
	w.Write([]byte("Name: " + name + "\n"))
	w.Write([]byte("Address: " + address + "\n"))
}
