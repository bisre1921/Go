package main

import (
	"fmt"
	"go-and-mysql/pkg/config"
	"go-and-mysql/pkg/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()
	router := routes.SetBookStoreRoutes()
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
