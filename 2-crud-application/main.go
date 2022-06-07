package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie = []Movie{}

func validateId(w http.ResponseWriter, r *http.Request) (string, error) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, isExists := params["id"]

	if !isExists {
		return "", fmt.Errorf("id not found in request")
	}

	return id, nil
}

func isMoviePresent(id string) (int, error) {

	for index, movie := range movies {
		if movie.ID == id {
			return index, nil
		}
	}

	return -1, fmt.Errorf("Movie not found")
}

func getAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}
func getMovies(w http.ResponseWriter, r *http.Request) {

	id, err := validateId(w, r)

	if err != nil {
		http.Error(w, "invalid request param id", http.StatusBadRequest)
	}

	index, err := isMoviePresent(id)

	if err != nil {
		http.Error(w, "data not present", http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(movies[index])

}
func createMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movies = append(movies, movie)

	movie.ID = strconv.Itoa(rand.Intn(1e10))

	json.NewEncoder(w).Encode(movie)

}
func updateMovies(w http.ResponseWriter, r *http.Request) {

	deleteMovies(w, r)

	createMovies(w, r)

}
func deleteMovies(w http.ResponseWriter, r *http.Request) {

	id, err := validateId(w, r)

	if err != nil {
		http.Error(w, "invalid request param id", http.StatusBadRequest)
	}

	index, err := isMoviePresent(id)

	if err != nil {
		http.Error(w, "data not present", http.StatusNotFound)
	}

	movies = append(movies[:index], movies[index:]...)

	json.NewEncoder(w).Encode(`{
		message : "Movie Deleted",
		status : 200
	}`)

}

func loggerMiddlerwere(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})

}

func main() {

	r := mux.NewRouter()

	r.Use(loggerMiddlerwere)

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("[Main] Starting server on PORT : %s", PORT)

	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatalln("[Main] Error starting server", err)
	}
}
