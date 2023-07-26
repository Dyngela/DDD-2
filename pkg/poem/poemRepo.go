package poem

import "gorm.io/gorm"

type IPoemRepo interface {
	Create(poem Poem) error
	FindById(int) (Poem, error)
}

type repository struct {
	db *gorm.DB
}

func NewPoemRepo(db *gorm.DB) IPoemRepo {
	return &repository{db: db}
}

func (r repository) Create(poem Poem) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindById(id int) (Poem, error) {
	//TODO implement me
	panic("implement me")
}
