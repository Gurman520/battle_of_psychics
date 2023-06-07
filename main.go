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
	router.HandleFunc("/result", server.ResultBattleHandler).Methods("POST")

	log.Println("Server is listening in 0.0.0.0:8080...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))
}
