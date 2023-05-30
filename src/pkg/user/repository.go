package user

import (
	"gorm.io/gorm"
)

type IUserRepo interface {
	CreateUser()
	UpdateUser()
	GetAllUser() ([]Users, error)
	GetUserById(uint) (Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) IUserRepo {
	return &repository{db: db}
}

func (r *repository) CreateUser() {
	//TODO implement me
	panic("implement me")
}

func (r *repository) UpdateUser() {
	//TODO implement me
	panic("implement me")
}

func (r *repository) GetAllUser() ([]Users, error) {
	var users []Users
	err := r.db.Model(&Users{}).Preload("Roles").Find(&users).Error
	return users, err
}

func (r *repository) GetUserById(u uint) (Users, error) {
	//TODO implement me
	panic("implement me")
}
