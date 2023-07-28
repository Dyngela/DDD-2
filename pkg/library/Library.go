package library

type Library struct {
	Books []Book
}

type Book struct {
	ID         int
	Author     string
	Date       string
	Genre      string
	IsBorrowed bool
}
