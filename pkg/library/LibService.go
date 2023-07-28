package library

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var libRepo ILibRepo

func InitLibService(db *gorm.DB) {
	libRepo = NewBookRepo(db)
}

type ILibService interface {
	Borrow(int, *gin.Context) error
	FindAll(ctx *gin.Context) ([]Book, error)
	FindById(i int, ctx *gin.Context) (Book, error)
	Delete(i int, ctx *gin.Context) error
	Create(book Book, ctx *gin.Context) error
}

type service struct{}

func NewLibService() ILibService {
	return &service{}
}

func (s service) Borrow(i int, context *gin.Context) error {
	isBorrow, err := libRepo.Borrow(i)
	if err != nil {
		return err
	}
	if isBorrow {
		return errors.New("Book is not available")
	} else {
		return nil
	}
}

func (s service) FindAll(ctx *gin.Context) ([]Book, error) {
	books, err := libRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s service) FindById(i int, ctx *gin.Context) (Book, error) {
	books, err := libRepo.FindById(i)
	if err != nil {
		return Book{}, err
	}
	return books, nil
}

func (s service) Delete(i int, ctx *gin.Context) error {
	err := libRepo.Delete(i)
	if err != nil {
		return err
	}
	return nil
}

func (s service) Create(book Book, ctx *gin.Context) error {
	err := libRepo.Create(book)
	if err != nil {
		return err
	}
	return nil
}
