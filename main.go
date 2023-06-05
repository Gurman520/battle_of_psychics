package main

import (
	"log"
	"net/http"

	back "battle_of_psychics/backend"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	server := back.CreateNewServer()

	router.HandleFunc("/startGame", server.StartGame).Methods("GET")
	router.HandleFunc("/hypotheses", server.Hypotheses).Methods("GET")
	router.HandleFunc("/rank", server.RankPsychics).Methods("POST")

	log.Println("Server is listening...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
