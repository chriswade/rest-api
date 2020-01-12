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
	user     = "chriswade"
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

//Add book to books table
func Add(book data.Book) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	sqlStatement := `
	WITH stmt1 AS (
		INSERT INTO books (isbn, title, author_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	)
		INSERT INTO author(firstname, lastname)
		VALUES ($6, $7)
		RETURNING id
`

	err = db.QueryRow(sqlStatement,
		book.ID,
		book.Isbn,
		book.Title,
		book.AuthorID,
		book.AuthorID,
		book.Author.Firstname,
		book.Author.Lastname).Scan(&book.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Book added:", book.ID)
}
