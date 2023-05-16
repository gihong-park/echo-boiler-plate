package dto

import (
	"blog_api/app/model"

	"gorm.io/gorm"
)

type UserRequest struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (userReq *UserRequest) ToModel() *model.User {
	return &model.User{
		Model:    gorm.Model{ID: userReq.ID},
		Name:     userReq.Name,
		Password: userReq.Password,
		Email:    userReq.Email,
	}
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
