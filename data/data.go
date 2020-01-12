package data

//Book Struct (Model/class/interface)
type Book struct {
	ID       int     `json:"id"`
	Isbn     int     `json:"isbn"`
	Title    string  `json:"title"`
	AuthorID int     `json:"authorid"`
	Author   *Author `json:"author"`
}

//Author stuct
type Author struct {
	ID        int    `json:id`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Init books var as a slick Book struct
var Books []Book
