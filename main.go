package main

import (
	"log"
	"net/http"

	svr "battle_of_psychics/backend/server"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	server := svr.CreateNewServer()

	// Handlers
	router.HandleFunc("/startGame", server.StartGame).Methods("GET")
	router.HandleFunc("/hypotheses", server.Hypotheses).Methods("GET")
	router.HandleFunc("/rank", server.RankPsychics).Methods("POST")

	log.Println("Server is listening in localhost:8080...")
	log.Println("Start Game in localhost:8080/startGame")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
