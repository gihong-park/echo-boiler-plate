package authService

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"

	"gorm.io/gorm"
)

type AuthService interface {
	Save(userReq *dto.UserRequest) (*model.User, error)
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	UpdateByID(userReq *dto.UserRequest) (*model.User, error)
	SetDB(db *gorm.DB)
}
