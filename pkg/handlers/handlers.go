package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World")
	if err != nil {
		fmt.Printf("Err: %+v", err)
	}
}

func HandleRequests(router *mux.Router) {
	router.HandleFunc("/", homePage).Methods("GET")
}