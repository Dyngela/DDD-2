package user

import (
	"awesomeProject/generic"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
	"log"
)

var userRepo IUserRepo
var genericRepo generic.IGenericRepo
var logger zerolog.Logger

func InitUserService(db *gorm.DB, log zerolog.Logger) {
	userRepo = NewUserRepo(db, log)
	genericRepo = generic.NewGenericRepo(db)
	logger = log
}

type IUserService interface {
	GetById()
	GetAll()
	Create()
	Update()
	Delete()
}

type service struct{}

func NewUserService() IUserService {
	return &service{}
}

func (u *service) GetById() {
	user, err := userRepo.FindById(0)
	if err != nil {
		panic(err)
	}
	logger.Info().Msg("user")
	log.Println(user)
}

func (u *service) GetAll() {
	//id, err := genericRepo.FindAll(Users{}, nil)
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("CA MARCHE")
	//log.Println(id)
}

func (u *service) Update() {
}

func (u *service) Create() {
}

func (u *service) Delete() {
}
