package user

import (
	"awesomeProject/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type IUserRepo interface {
	Create(users Users) error
	Update(users Users) error
	Delete(id uint) error
	FindAll(*gin.Context) ([]Users, error)
	FindById(id uint) (Users, error)
}

type repository struct {
	db     *gorm.DB
	logger zerolog.Logger
}

func NewUserRepo(db *gorm.DB, logger zerolog.Logger) IUserRepo {
	return &repository{db: db, logger: logger}
}

func (r *repository) Create(users Users) error {
	err := r.db.Create(&users).Error
	if err != nil {
		return err
	}
	r.logger.Info().Msg("User created")
	return err
}

func (r *repository) Update(users Users) error {
	err := r.db.Save(&users).Error
	if err != nil {
		return err
	}
	r.logger.Info().Msg("User updated")
	return err
}

func (r *repository) FindAll(ctx *gin.Context) ([]Users, error) {
	var users []Users
	err := r.db.Model(&Users{}).Preload("Roles").Scopes(utils.Paginate(ctx.Request)).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}

func (r *repository) FindById(id uint) (Users, error) {
	var user Users
	err := r.db.Model(&Users{}).Preload("Roles").Find(&user, id).Error
	if err != nil {
		return Users{}, err
	}
	return user, nil
}

func (r *repository) Delete(id uint) error {
	err := r.db.Delete(&Users{}, id).Error
	r.logger.Info().Msg("User deleted")
	return err
}
