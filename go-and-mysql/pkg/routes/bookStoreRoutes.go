package routes

import (
	"go-and-mysql/pkg/controllers"

	"github.com/gorilla/mux"
)

// SetBookStoreRoutes sets the routes for the book store
func SetBookStoreRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
	return router
}
