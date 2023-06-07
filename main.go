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
	router.HandleFunc("/", server.Home).Methods("GET")
	router.HandleFunc("/startGame", server.StartHandler).Methods("GET")
	router.HandleFunc("/hypotheses", server.HypothesesHandler).Methods("GET")
	router.HandleFunc("/rank", server.RankPsychicsHandler).Methods("POST")

	log.Println("Server is listening in 127.0.0.1:8080...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
