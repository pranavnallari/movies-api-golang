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

// Movie struct represents a movie object
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director struct represents the director of a movie
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

// movies_seed is a slice to store the initial movie data
var movies_seed []Movie

// getMovies handles GET requests to fetch all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies_seed)
}

// deleteMovie handles DELETE requests to delete a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies_seed {
		if item.ID == params["id"] {
			movies_seed = append(movies_seed[:index], movies_seed[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies_seed)
}

// getMovie handles GET requests to fetch a movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies_seed {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createMovie handles POST requests to create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies_seed = append(movies_seed, movie)
	json.NewEncoder(w).Encode(movie)
}

// updateMovie handles PUT requests to update a movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies_seed {
		if item.ID == params["id"] {
			movies_seed = append(movies_seed[:index], movies_seed[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies_seed = append(movies_seed, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main() {
	r := mux.NewRouter()

	// Initializing the movie data
	movies_seed = append(movies_seed, Movie{ID: "1", Isbn: "4188839", Title: "Sample Movie Title 1", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies_seed = append(movies_seed, Movie{ID: "2", Isbn: "5633522", Title: "Sample Movie Title 2", Director: &Director{FirstName: "Mark", LastName: "Hill"}})
	movies_seed = append(movies_seed, Movie{ID: "3", Isbn: "7856465", Title: "Sample Movie Title 3", Director: &Director{FirstName: "Dash", LastName: "Breck"}})
	movies_seed = append(movies_seed, Movie{ID: "4", Isbn: "1234567", Title: "Sample Movie Title 4", Director: &Director{FirstName: "Peter", LastName: "Chris"}})

	// Define API endpoints
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Start server on port 8080
	fmt.Printf("Starting Server at Port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
