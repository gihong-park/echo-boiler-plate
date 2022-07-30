package todoRepository

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Save(todo *dto.TodoDTO) (*model.Todo, error)
	GetByID(id uint) (*model.Todo, error)
	GetAll() (todo *[]model.Todo, err error)
	UpdateByID(todo *dto.TodoDTO) (*model.Todo, error)
	SetDB(DB *gorm.DB)
	GetDB() *gorm.DB
}
