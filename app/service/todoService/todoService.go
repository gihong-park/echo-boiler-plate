package todoService

import (
	"blog_api/app/controller/dto"
	"blog_api/app/db"
	"blog_api/app/model"
	"blog_api/app/repository/todoRepository"
)

type TodoService interface {
	Save(todoDTO *dto.TodoDTO) model.Todo
}

func NewTodoService(todoRepo todoRepository.TodoRepository) TodoService {
	database := db.GetDB("sqlite")
	return &TodoServiceImpl{todoRepo: todoRepository.NewTodoRepository(database)}
}
