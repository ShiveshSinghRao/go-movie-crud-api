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

type Movie struct{ 
	ID         	string   `json:"id"`   //this is back-tick not single quotes
	Isbn        string 	`json:"isbn"`
    Title       string 	`json:"title"`
	Director *Director 	`json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index, item:=range movies{ 
		if item.ID==params["id"]{
         movies= append(movies[:index],movies[index+1:]...)
			break
		}
    json.NewEncoder(w).Encode(movies)
		
     }
}
func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params:=mux.Vars(r)
	for _,item:=range movies{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return 
		}
	}
}
func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(1000000))
	movies=append(movies,movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	var updatedMovie Movie
	err := json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for index, movie := range movies {
		if movie.ID == id {
			// Update the specific fields if they are provided in the request body
			if updatedMovie.Isbn != "" {
				movies[index].Isbn = updatedMovie.Isbn
			}
			if updatedMovie.Title != "" {
				movies[index].Title = updatedMovie.Title
			}
			if updatedMovie.Director != nil {
				if updatedMovie.Director.Firstname != "" {
					movies[index].Director.Firstname = updatedMovie.Director.Firstname
				}
				if updatedMovie.Director.Lastname != "" {
					movies[index].Director.Lastname = updatedMovie.Director.Lastname
				}
			}

			json.NewEncoder(w).Encode(movies[index])
			return
		}
	}

	http.Error(w, "Movie not found", http.StatusNotFound)
}


func main(){
	r:=mux.NewRouter()
	movies=append(movies,Movie{ID:"1",Isbn:"4354353",Title:"Movie one",Director: &Director{Firstname:"jhon",Lastname:"wick"}})
	movies=append(movies,Movie{ID:"2",Isbn:"34534645",Title:"Movie two",Director: &Director{Firstname:"papi",Lastname:"boss"}})
	movies= append(movies,Movie{ID:"3",Isbn:"24545234",Title:"Movie three",Director:&Director{Firstname:"Shree",Lastname:"Ram"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server ar port 8091\n")
	log.Fatal(http.ListenAndServe(":8091",r))
}
