package todoRepository

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"
	"log"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	DB *gorm.DB
}

func (repo *TodoRepositoryImpl) Save(todoDTO *dto.TodoDTO) model.Todo {
	todo := todoDTO.ToModel()
	result := repo.DB.Create(&todo)
	log.Default().Printf("affected rows are %v", result.RowsAffected)
	return *todo
}

func (repo *TodoRepositoryImpl) SetDB(DB *gorm.DB) {
	repo.DB = DB
}

func (repo *TodoRepositoryImpl) Init() TodoRepository {
	return &TodoRepositoryImpl{}
}
