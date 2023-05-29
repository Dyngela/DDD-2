package user

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name    string /*`gorm:"column:field_a"`*/
	Roles   Roles
	RolesID uint
}

func (u *Users) Validate() error {
	return nil
}

type Roles struct {
	gorm.Model
	RoleName string
	Users    []Users `gorm:"foreignKey:roles_id"`
}
