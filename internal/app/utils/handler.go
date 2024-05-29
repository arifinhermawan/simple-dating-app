package utils

import (
	"log"
	"net/http"

	"github.com/arifinhermawan/simple-dating-app/internal/app/server"
	"github.com/gorilla/mux"
)

func HandleRequest(handlers *server.Handler) {
	router := mux.NewRouter().StrictSlash(true)

	handleGetRequest(handlers, router)
	handlePostRequest(handlers, router)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleGetRequest(handlers *server.Handler, router *mux.Router) {
}

func handlePostRequest(handlers *server.Handler, router *mux.Router) {
}
