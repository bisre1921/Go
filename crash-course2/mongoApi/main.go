package main

import (
	"fmt"
	"log"
	"mongoApi/routers"
	"net/http"
)

func main() {
	fmt.Println("Mongo db api")

	r := routers.Router()
	fmt.Println("server is running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))

}
