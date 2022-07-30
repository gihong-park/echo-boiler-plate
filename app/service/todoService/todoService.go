package todoService

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"
	tr "blog_api/app/repository/todoRepository"
)

type TodoService interface {
	Save(todoDTO *dto.TodoDTO) (*model.Todo, error)
	GetByID(id uint) (*model.Todo, error)
	GetAll() (*[]model.Todo, error)
	UpdateByID(todoDTO *dto.TodoDTO) (*model.Todo, error)
	SetRepo(todoRepo tr.TodoRepository) tr.TodoRepository
}
