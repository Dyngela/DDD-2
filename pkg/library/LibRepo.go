package library

import "gorm.io/gorm"

type ILibRepo interface {
	Borrow(int) (bool, error)
	FindAll() ([]Book, error)
	FindById(int) (Book, error)
	Delete(int) error
	Create(Book) error
}

type repo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) ILibRepo {
	return &repo{db: db}
}

func (r repo) Borrow(i int) (bool, error) {
	// find the book and check if it is borrowed or not then send the bool accordingly
	return true, nil
}

func (r repo) FindAll() ([]Book, error) {
	return nil, nil
}

func (r repo) FindById(i int) (Book, error) {
	return Book{}, nil
}

func (r repo) Delete(i int) error {
	return nil
}

func (r repo) Create(book Book) error {
	return nil
}
