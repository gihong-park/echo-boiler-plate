package todoRepository

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"
	"log"
	"time"

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

func (repo *TodoRepositoryImpl) GetAll() (todos *[]model.Todo, err error) {
	err = repo.DB.Debug().Model(&model.Todo{}).Find(&todos).Error
	if err != nil {
		return todos, err
	}
	return todos, nil
}

func (repo *TodoRepositoryImpl) UpdateByID(todoDTO *dto.TodoDTO) (todo *model.Todo, err error) {
	todo, err = repo.GetByID(todoDTO.ID)
	if err != nil {
		return nil, err
	}

	todo.Body = todoDTO.Body
	todo.UpdatedAt = time.Now().UTC()

	err = repo.DB.Debug().Model(model.Todo{}).Where("id = ?", todoDTO.ID).Updates(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (repo *TodoRepositoryImpl) SetDB(DB *gorm.DB) {
	repo.DB = DB
}

func (repo *TodoRepositoryImpl) GetDB() *gorm.DB {
	return repo.DB
}
