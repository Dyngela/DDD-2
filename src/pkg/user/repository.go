package user

import (
	"gorm.io/gorm"
)

type userInterface interface {
	CreateUser()
	UpdateUser()
}

type UserRepo struct {
	repo *gorm.DB
	userInterface
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{repo: db}
}

func (u *UserRepo) CreateUser() {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) UpdateUser() {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) GetAllUser() ([]Users, error) {
	var users []Users
	err := u.repo.Model(&Users{}).Preload("Roles").Find(&users).Error
	return users, err
}
