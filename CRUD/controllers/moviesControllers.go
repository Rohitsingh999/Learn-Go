package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    ` json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director" `
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func init() {

	movies = append(movies, Movie{

		ID:    "1",
		Isbn:  "42342",
		Title: "Kung-fu-panda",
		Director: &Director{
			Firstname: "Rohit",
			Lastname:  "Rawat",
		},
	})
}

//Get all movies

func GetMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(movies)
}

// Delete movie based on ID
func DeleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, items := range movies {

		if items.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	response := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Movie with id %s deleted successfully", params["id"]),
	}

	json.NewEncoder(w).Encode(response)

}

// Get movie based on ID ..
func GetMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	response := make(map[string]interface{})

	for _, item := range movies {

		if item.ID == params["id"] {

			response["success"] = true
			response["message"] = item
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	response["success"] = false
	response["message"] = fmt.Sprintf("data with %s is not present", params["id"])
	json.NewEncoder(w).Encode(response)

}

// create movie
func CreateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
	}

	movie.ID = strconv.Itoa(rand.Intn(100000000))

	movies = append(movies, movie)

	response := map[string]interface{}{
		"success": true,
		"message": "Movie created successfully",
		"data":    movie,
	}
	json.NewEncoder(w).Encode(response)

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application-json")

	params := mux.Vars(r)

	var updatemovie Movie

	json.NewDecoder(r.Body).Decode(&updatemovie)

	for index, item := range movies {

		if item.ID == params["id"] {

			movies[index] = updatemovie
			movies[index].ID = params["id"]

			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	response := map[string]string{

		"message": "Not Update",
	}
	json.NewEncoder(w).Encode(response)

}
