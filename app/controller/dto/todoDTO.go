package dto

import (
	"blog_api/app/model"

	"gorm.io/gorm"
)

type TodoDTO struct {
	ID   uint   `json:"id"`
	Body string `json:"body"`
}

func (todoDTO *TodoDTO) ToModel() *model.Todo {
	return &model.Todo{Model: gorm.Model{ID: todoDTO.ID}, Body: todoDTO.Body}
}
