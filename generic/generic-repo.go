package generic

import (
	"gorm.io/gorm"
)

type IGenericRepo interface {
	Create(T any) error
	Update(T any) error
	Delete(T any, id uint) error
	//FindAll(T any, ctx *gin.Context) ([]any, error)
	//FindById(T any, id uint) (any, error)
}

type repository struct {
	db *gorm.DB
}

func NewGenericRepo(db *gorm.DB) IGenericRepo {
	return &repository{db: db}
}

func (r *repository) Create(T any) error {
	err := r.db.Create(&T).Error
	if err != nil {
		return err
	}
	return err
}

func (r *repository) Update(T any) error {
	err := r.db.Save(&T).Error
	if err != nil {
		return err
	}
	return err
}

func (r *repository) Delete(T any, id uint) error {
	err := r.db.Delete(&T, id).Error
	return err
}

//func (r *repository) FindAll(T any, ctx *gin.Context) ([]any, error) {
//	var generic []any
//
//	err := r.db.Model(&T).Scopes(common.Paginate(ctx.Request)).Find(&generic).Error
//	if err != nil {
//		return nil, err
//	}
//	return generic, err
//}
//
//func (r *repository) FindById(T any, id uint) (any, error) {
//	generic := reflect.ValueOf(T)
//	err := r.db.Model(&T).Find(&generic, id).Error
//	if err != nil {
//		return T, err
//	}
//	return generic, nil
//}
