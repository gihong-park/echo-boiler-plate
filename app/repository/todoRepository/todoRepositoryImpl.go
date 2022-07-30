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

func (repo *TodoRepositoryImpl) Save(todoDTO *dto.TodoDTO) (todo *model.Todo, err error) {
	todo = todoDTO.ToModel()
	err = repo.DB.Debug().Model(&model.Todo{}).Create(&todo).Error
	if err != nil {
		log.Println(err)
		return &model.Todo{}, err
	}

	return todo, nil
}

func (repo *TodoRepositoryImpl) GetByID(id uint) (todo *model.Todo, err error) {
	err = repo.DB.Debug().Model(&model.Todo{}).First(&todo).Error
	if err != nil {
		return &model.Todo{}, err
	}
	return todo, nil
}

func (repo *TodoRepositoryImpl) GetAll() (todo *[]model.Todo, err error) {
	var todos []model.Todo
	err = repo.DB.Debug().Model(&model.Todo{}).Find(&todos).Error
	if err != nil {
		return &todos, err
	}
	return &todos, nil
}

func (repo *TodoRepositoryImpl) SetDB(DB *gorm.DB) {
	repo.DB = DB
}

func (repo *TodoRepositoryImpl) GetDB() *gorm.DB {
	return repo.DB
}
