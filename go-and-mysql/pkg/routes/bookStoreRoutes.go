package routes

import (
	"go-and-mysql/pkg/controllers"

	"github.com/gorilla/mux"
)

// SetBookStoreRoutes sets the routes for the book store API
func SetBookStoreRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/bookstore", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/api/bookstore", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/bookstore/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/bookstore/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/bookstore/{id}", controllers.DeleteBook).Methods("DELETE")
	return router
}
