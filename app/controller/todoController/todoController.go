package todoController

import (
	"blog_api/app/controller/dto"
	ts "blog_api/app/service/todoService"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type TodoController struct {
	todoServ ts.TodoService
}

func (todoController *TodoController) SaveHandler(c echo.Context) error {
	t := new(dto.TodoDTO)
	c.Bind(t)
	todo, err := todoController.todoServ.Save(t)
	if err != nil {
		log.Errorf("fail to save todo: %+v", err)
		return err
	}
	return c.JSON(http.StatusCreated, todo)
}
