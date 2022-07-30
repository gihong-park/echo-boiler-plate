package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	name       string
	nickname   string
	password   string
	email      string
	emailCheck bool
}
