package todoController

import (
	"blog_api/app/repository/todoRepository"
	"blog_api/app/service/todoService"

	"gorm.io/gorm"
)

func NewTodoController(s todoService.TodoService) TodoController {
	return TodoController{s}
}

func NewTodoService(r todoRepository.TodoRepository) todoService.TodoService {
	todoServ := &todoService.TodoServiceImpl{}
	todoServ.SetRepo(r)
	return todoServ
}

func NewTodoRepository(db *gorm.DB) todoRepository.TodoRepository {
	todoRepo := &todoRepository.TodoRepositoryImpl{}
	todoRepo.SetDB(db)
	return todoRepo
}
