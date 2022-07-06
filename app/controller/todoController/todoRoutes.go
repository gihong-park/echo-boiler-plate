package todoController

import "github.com/labstack/echo/v4"

func (todoController *TodoController) Routes(g *echo.Group) {
	//api/v1/todo
	g.POST("", todoController.SaveHandler)
}
