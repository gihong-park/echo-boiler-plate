package todoService

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"
	"blog_api/app/repository/todoRepository"
	tr "blog_api/app/repository/todoRepository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoServiceImpl struct {
	todoRepo todoRepository.TodoRepository
}

func (todoServ *TodoServiceImpl) SetRepo(todoRepo tr.TodoRepository) todoRepository.TodoRepository {
	todoServ.todoRepo = todoRepo
	return todoServ.todoRepo
}

func (todoServ *TodoServiceImpl) Save(todoDTO *dto.TodoDTO) (*model.Todo, error) {
	if todoDTO.Body == "" {
		return &model.Todo{}, echo.NewHTTPError(http.StatusBadRequest, "todo must not be empty")
	}
	return todoServ.todoRepo.Save(todoDTO)
}

func (todoServ *TodoServiceImpl) GetByID(id uint) (*model.Todo, error) {
	return todoServ.todoRepo.GetByID(id)
}

func (todoServ *TodoServiceImpl) GetAll() (*[]model.Todo, error) {
	return todoServ.todoRepo.GetAll()
}

func (todoServ *TodoServiceImpl) UpdateByID(todoDTO *dto.TodoDTO) (*model.Todo, error) {
	return todoServ.todoRepo.UpdateByID(todoDTO)
}
