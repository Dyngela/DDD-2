package poem

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

var poemRepo IPoemRepo

func InitPoemService(db *gorm.DB) {
	poemRepo = NewPoemRepo(db)
}

type IPoemService interface {
	GetById(ctx *gin.Context) (Poem, error)
	Create(ctx *gin.Context) error
}

type service struct{}

func NewPoemService() IPoemService {
	return &service{}
}

func (s service) GetById(ctx *gin.Context) (Poem, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	poem, err := poemRepo.FindById(id)
	if err != nil {
		return Poem{}, err
	}
	return poem, err
}

func (s service) Create(ctx *gin.Context) error {
	var request CreatePoemRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		return err
	}
	err = poemRepo.Create(Poem{
		Author:      request.Author,
		ReleaseDate: request.ReleaseDate,
		Content:     request.Content,
	})
	if err != nil {
		return err
	}
	return nil
}
