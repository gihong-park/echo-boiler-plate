package todoService

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"
	tr "blog_api/app/repository/todoRepository"
)

type TodoService interface {
	Init() TodoService
	Save(todoDTO *dto.TodoDTO) (*model.Todo, error)
	SetRepo(todoRepo tr.TodoRepository) tr.TodoRepository
}

func NewTodoService[T TodoService]() TodoService {
	var todoServ T
	return todoServ.Init()
}
