package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

var userRepo IUserRepo

func InitUserService(db *gorm.DB) {
	userRepo = NewUserRepo(db)
}

type IUserService interface {
	GetById(*gin.Context)
	GetAll(*gin.Context)
}

type service struct{}

func NewUserService() IUserService {
	return &service{}
}

func (u *service) GetById(context *gin.Context) {
	user, err := userRepo.GetUserById(0)
	if err != nil {
		panic(err)
	}
	log.Println(user)
}

func (u *service) GetAll(context *gin.Context) {
}
