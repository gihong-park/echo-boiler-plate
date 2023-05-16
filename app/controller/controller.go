package controller

import (
	"blog_api/app/controller/authController"
	"blog_api/app/controller/indexController"
	"blog_api/app/controller/todoController"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	TodoCont  todoController.TodoController
	IndexCont indexController.IndexController
	AuthCont  authController.AuthController
}

func (controller *Controller) Routes(echoServer *echo.Echo) {
	authGroup := echoServer.Group("api/v1/auth")
	todoGroup := echoServer.Group("api/v1/todo")
	indexGroup := echoServer.Group("api/v1")
	healthGroup := echoServer.Group("")

	controller.TodoCont.Routes(todoGroup)
	controller.IndexCont.Routes(indexGroup)
	controller.IndexCont.Routes(healthGroup)
	controller.AuthCont.Routes(authGroup)
}
