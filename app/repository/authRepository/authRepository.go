package authRepository

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Init() AuthRepository
	Save(todo *dto.TodoDTO) model.Todo
	SetDB(DB *gorm.DB)
}

func NewTodoRepository[T AuthRepository]() AuthRepository {
	var todoRepo T
	return todoRepo.Init()
}
