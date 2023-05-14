package todoRepository

import (
	"blog_api/app/controller/dto"
	"blog_api/app/model"
	"fmt"
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
		return &model.Todo{}, fmt.Errorf("repository create todo error: %w", err)
	}

	return todo, nil
}

func (repo *TodoRepositoryImpl) GetByID(id uint) (todo *model.Todo, err error) {
	err = repo.DB.Debug().Model(&model.Todo{}).First(&todo).Error
	if err != nil {
		return &model.Todo{}, fmt.Errorf("repository get todo by id error: %w", err)
	}
	return todo, nil
}

func (repo *TodoRepositoryImpl) GetAll() (todos *[]model.Todo, err error) {
	err = repo.DB.Debug().Model(&model.Todo{}).Find(&todos).Error
	if err != nil {
		return todos, fmt.Errorf("repository get all todo error: %w", err)
	}
	return todos, nil
}

func (repo *TodoRepositoryImpl) UpdateByID(todoDTO *dto.TodoDTO) (todo *model.Todo, err error) {
	todo, err = repo.GetByID(todoDTO.ID)
	if err != nil {
		return nil, fmt.Errorf("repository fail to get todo by id: %w", err)
	}

	todo.Body = todoDTO.Body
	todo.UpdatedAt = time.Now().UTC()

	err = repo.DB.Debug().Model(model.Todo{}).Where("id = ?", todoDTO.ID).Updates(&todo).Error
	if err != nil {
		return nil, fmt.Errorf("repository fail to update todo: %w", err)
	}
	return todo, nil
}

func (repo *TodoRepositoryImpl) SetDB(DB *gorm.DB) {
	repo.DB = DB
}

func (repo *TodoRepositoryImpl) GetDB() *gorm.DB {
	return repo.DB
}
