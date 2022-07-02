package dto

import (
	"blog_api/app/model"

	"gorm.io/gorm"
)

type TodoDTO struct {
	Body string
}

func (todoDTO *TodoDTO) ToModel() *model.Todo {
	return &model.Todo{Model: gorm.Model{}, Body: todoDTO.Body}
}
