//go:build wireinject
// +build wireinject

package todoController

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitTodoController(db *gorm.DB) TodoController {
	wire.Build(NewTodoController, NewTodoService, NewTodoRepository)

	return TodoController{}
}
