package main

import (
	"log"
	"net/http"

	"github.com/chriswade/rest-api/data"
	"github.com/chriswade/rest-api/router"
	"github.com/gorilla/mux"
)

func main() {
	// Init the mux router
	r := mux.NewRouter()

	//mock data
	data.Books = append(data.Books, data.Book{ID: "1", Isbn: "123", Title: "book1", Author: &data.Author{Firstname: "chris", Lastname: "wade"}})
	data.Books = append(data.Books, data.Book{ID: "2", Isbn: "124", Title: "book2", Author: &data.Author{Firstname: "natalie", Lastname: "hammond"}})

	// Create route handers
	r.HandleFunc("/api/books", router.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", router.GetBook).Methods("GET")
	r.HandleFunc("/api/books", router.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", router.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", router.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
