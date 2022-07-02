package todoRepository

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Save(todo *dto.TodoDTO) model.Todo
}

func NewTodoRepository(DB *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{DB}
}
