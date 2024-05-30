package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arifinhermawan/simple-dating-app/internal/app/server"
	"github.com/gorilla/mux"
)

func HandleRequest(handlers *server.Handler) {
	router := mux.NewRouter().StrictSlash(true)

	// handleGetRequest(router)
	handlePostRequest(handlers, router)

	log.Println("Serving at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleGetRequest(router *mux.Router) {
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Hello world")
	}).Methods("GET")
}

func handlePostRequest(handlers *server.Handler, router *mux.Router) {
	router.HandleFunc("/account/signup", handlers.Account.HandlerCreateUserAccount).Methods("POST")
}
