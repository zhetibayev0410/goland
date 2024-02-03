package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Response struct {
	Clubs []Club `json:"clubs"`
}

type Club struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Coach      string `json:"coach"`
	Founded    int    `json:"founded"`
	Stadium    string `json:"stadium"`
	HomeColors string `json:"home_colors"`
	AwayColors string `json:"away_colors"`
}

func main() {
	log.Println("Starting Football API server")
	// Create a new router
	router := mux.NewRouter()
	log.Println("Creating routes")
	// Specify endpoints
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/clubs", Clubs).Methods("GET")
	http.Handle("/", router)

	// Start and listen to requests
	http.ListenAndServe("localhost:8080", router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering health check endpoint")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Football API is up and running")
}

func Clubs(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering clubs endpoint")
	var response Response
	clubs := prepareResponse()

	response.Clubs = clubs

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}

func prepareResponse() []Club {
	var clubs []Club

	var club Club
	club.Id = 1
	club.Name = "Liverpool FC"
	club.Coach = "Jurgen Klopp"
	club.Founded = 1892
	club.Stadium = "Anfield Road"
	club.HomeColors = "Red"
	club.AwayColors = "Purple"
	clubs = append(clubs, club)

	club.Id = 2
	club.Name = "Arsenal"
	club.Coach = " Mikel Arteta"
	club.Founded = 1886
	club.Stadium = "Emirates Stadium"
	club.HomeColors = "Red"
	club.AwayColors = "Green"
	clubs = append(clubs, club)

	return clubs
}
