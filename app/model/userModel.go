package model

import (
	"blog_api/app/auth/role"
	"blog_api/app/security"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Email    string `gorm:"unique"`
	Role     role.Role
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}
