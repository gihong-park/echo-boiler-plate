package controller

import (
	"blog_api/app/controller/authController"
	"blog_api/app/controller/indexController"
	"blog_api/app/controller/todoController"

	"gorm.io/gorm"
)

func NewController(db *gorm.DB) Controller {
	todoCont := todoController.InitTodoController(db)
	authCont := authController.InitAuthController(db)
	return Controller{todoCont, indexController.IndexController{}, authCont}
}
