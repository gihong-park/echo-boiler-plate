package todoService

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"
	"blog_api/app/repository/todoRepository"
	tr "blog_api/app/repository/todoRepository"
)

type TodoServiceImpl struct {
	todoRepo todoRepository.TodoRepository
}

func (todoServ *TodoServiceImpl) Init() TodoService {
	return &TodoServiceImpl{}
}

func (todoServ *TodoServiceImpl) SetRepo(todoRepo tr.TodoRepository) todoRepository.TodoRepository {
	todoServ.todoRepo = todoRepo
	return todoServ.todoRepo
}

func (todoServ *TodoServiceImpl) Save(todoDTO *dto.TodoDTO) (*model.Todo, error) {
	return todoServ.todoRepo.Save(todoDTO)
}
