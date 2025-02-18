package routers

import (
	"mongoApi/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controllers.DeleteAllMovies).Methods("DELETE")

	return router
}
