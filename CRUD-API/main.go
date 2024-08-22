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

type movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []movie

func getmovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["Id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
func deletemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.Id == params["Id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func createmovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var Movie movie
	_ = json.NewDecoder(r.Body).Decode(&Movie)
	Movie.Id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, Movie)
	json.NewEncoder(w).Encode(Movie)
}
func updatemovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("content-type ", "json/application")
	//params(request)
	params := mux.Vars(r)
	//loop over the movie, range
	for index, item := range movies {

		if item.Id == params["Id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["Id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, movie{Id: "1", Isbn: "2345", Title: "movie one", Director: &Director{Firstname: "john", Lastname: "akeal"}})
	movies = append(movies, movie{Id: "2", Isbn: "4345", Title: "movie two", Director: &Director{Firstname: "smith", Lastname: "akeal"}})
	movies = append(movies, movie{Id: "3", Isbn: "6345", Title: "movie three", Director: &Director{Firstname: "lial", Lastname: "akeal"}})

	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{Id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", createmovie).Methods("POST")
	r.HandleFunc("/movie/{Id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movie/{Id}", deletemovie).Methods("DELETE")

	fmt.Println("Starting server at 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
