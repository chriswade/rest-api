package router

import (
	"encoding/json"
	"net/http"

	"github.com/chriswade/rest-api/data"
	"github.com/chriswade/rest-api/database"
)

//Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Books)
}

//Get book
// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) //get params
// 	//loop through books
// 	for _, item := range data.Books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&data.Book{})
// }

//Create book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book data.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	database.Add(book)
	json.NewEncoder(w).Encode(book)

}

//Update book
// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range data.Books {
// 		if item.ID == params["id"] {
// 			data.Books = append(data.Books[:index], data.Books[index+1:]...)
// 			var book data.Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			data.Books = append(data.Books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(data.Books)
// }

//delete book
// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range data.Books {
// 		if item.ID == params["id"] {
// 			data.Books = append(data.Books[:index], data.Books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(data.Books)
// }
