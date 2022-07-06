package todoController

import (
	"blog_api/app/controller/dto"
	ts "blog_api/app/service/todoService"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	todoServ ts.TodoService
}

func (todoController *TodoController) SaveHandler(c echo.Context) error {
	t := new(dto.TodoDTO)
	c.Bind(t)
	return c.JSON(http.StatusCreated, todoController.todoServ.Save(t))
}
