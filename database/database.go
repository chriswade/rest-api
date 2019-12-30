package database

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

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

func Add(book data.Book) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	sqlStatement := `
	WITH stmt1 AS (
		INSERT INTO books (id, isbn, title, author_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	)
		INSERT INTO author(id, firstname, lastname)
		VALUES ($5, $6, $7)
		RETURNING id
`
	id := 0
	book.ID = strconv.Itoa(rand.Intn(10000000))
	book.AuthorID = strconv.Itoa(rand.Intn(10000000))
	authorid := 0
	err = db.QueryRow(sqlStatement,
		book.ID,
		book.Isbn,
		book.Title,
		book.AuthorID,
		book.AuthorID,
		book.Author.Firstname,
		book.Author.Lastname).Scan(&id)
	id = id + 1
	authorid = authorid + 1
	if err != nil {
		panic(err)
	}
	fmt.Println("Book added:", book.Title)
}
