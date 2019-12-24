package database

import (
	"database/sql"
	"fmt"

	"github.com/chriswade/rest-api/data"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "cwade"
	password = "password"
	dbname   = "demo"
)

func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func Add(book data.Book) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	sqlStatement := `
	INSERT INTO books (id, isbn, title, author)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, book.ID, book.Isbn, book.Title, book.Author).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Book added:", book.Title)
}
