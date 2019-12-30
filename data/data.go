package data

//Book Struct (Model/class/interface)
type Book struct {
	ID       string  `json:"id"`
	Isbn     string  `json:"isbn"`
	Title    string  `json:"title"`
	AuthorID string  `json:"authorid"`
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
