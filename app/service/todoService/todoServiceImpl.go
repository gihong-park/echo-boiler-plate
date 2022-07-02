package todoService

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"
	"blog_api/app/repository/todoRepository"
)

type TodoServiceImpl struct {
	todoRepo todoRepository.TodoRepository
}

func (todoServ *TodoServiceImpl) Save(todoDTO *dto.TodoDTO) model.Todo {
	return todoServ.todoRepo.Save(todoDTO)
}
