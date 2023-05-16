package todoController

import (
	"blog_api/app/auth/role"
	"blog_api/app/middlewares"

	"github.com/labstack/echo/v4"
)

func (todoController *TodoController) Routes(g *echo.Group) {
	//api/v1/todo
	g.POST("", todoController.SaveHandler, middlewares.TokenAuthMiddleware(role.Member))
}
