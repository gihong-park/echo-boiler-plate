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

func (todoServ *TodoServiceImpl) SetRepo(todoRepo tr.TodoRepository) todoRepository.TodoRepository {
	todoServ.todoRepo = todoRepo
	return todoServ.todoRepo
}

func (todoServ *TodoServiceImpl) Save(todoDTO *dto.TodoDTO) (*model.Todo, error) {
	return todoServ.todoRepo.Save(todoDTO)
}

func (todoServ *TodoServiceImpl) GetByID(id uint) (*model.Todo, error) {
	todo, err := todoServ.todoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (todoServ *TodoServiceImpl) GetAll() (*[]model.Todo, error) {
	todos, err := todoServ.todoRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (todoServ *TodoServiceImpl) UpdateByID(todoDTO *dto.TodoDTO) (*model.Todo, error) {
	todo, err := todoServ.todoRepo.UpdateByID(todoDTO)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
