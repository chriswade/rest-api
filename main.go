package main

import (
	"log"
	"net/http"

	"github.com/chriswade/rest-api/database"
	"github.com/chriswade/rest-api/router"
	"github.com/gorilla/mux"
)

func main() {
	// Init the mux router
	r := mux.NewRouter()

	database.Connect()

	// Create route handers
	r.HandleFunc("/api/books", router.GetBooks).Methods("GET")
	// r.HandleFunc("/api/books/{id}", router.GetBook).Methods("GET")
	r.HandleFunc("/api/books", router.CreateBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", router.UpdateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", router.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
