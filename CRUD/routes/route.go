package routes

import (
	"CRUD/controllers"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")

	return router

}
