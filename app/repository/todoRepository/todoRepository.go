package todoRepository

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Init() TodoRepository
	Save(todo *dto.TodoDTO) model.Todo
	SetDB(DB *gorm.DB)
}

func NewTodoRepository[T TodoRepository]() TodoRepository {
	var todoRepo T
	return todoRepo.Init()
}
