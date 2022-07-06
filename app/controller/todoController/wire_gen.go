// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package todoController

import (
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitTodoController(db *gorm.DB) TodoController {
	todoRepository := NewTodoRepository(db)
	todoService := NewTodoService(todoRepository)
	todoController := NewTodoController(todoService)
	return todoController
}
