package authRepository

import "gorm.io/gorm"

type AuthRepositoryImpl struct {
	DB *gorm.DB
}
