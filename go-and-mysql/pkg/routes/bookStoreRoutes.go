package routes

import (
	"github.com/gorilla/mux"
	"go-and-mysql/pkg/controllers"
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
