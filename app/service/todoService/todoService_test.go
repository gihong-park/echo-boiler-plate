package todoService

import (
	"blog_api/app/controller/dto"
	"blog_api/app/db"
	"blog_api/app/model"
	"blog_api/app/repository/todoRepository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	todoServ TodoService
	todoRepo todoRepository.TodoRepository
	todoDTO  dto.TodoDTO = dto.TodoDTO{Body: "save item"}
)

func init() {
	database = db.GetDB("sqlite")
	database.Migrator().DropTable(&model.Todo{})
	database.AutoMigrate(&model.Todo{})
	todoServ = &TodoServiceImpl{}
	todoRepo = &todoRepository.TodoRepositoryImpl{DB: database}
	todoServ.SetRepo(todoRepo)
}

func TestTodoService(t *testing.T) {

	assert.Equal(t, todoRepo, todoServ.SetRepo(todoRepo))
}

func TestTodoSave(t *testing.T) {
	todoModel, err := todoServ.Save(&todoDTO)
	if err != nil {
		t.Fatalf("TodoService has failed: %v", err)
	}

	t.Log(todoModel)
	assert.Equal(t, todoDTO.Body, todoModel.Body)

}

func TestTodoSaveBadRequest(t *testing.T) {
	dto := dto.TodoDTO{Body: ""}
	todoModel, err := todoServ.Save(&dto)
	assert.IsType(t, &model.Todo{}, todoModel)
	assert.ErrorContains(t, err, "code=")
}

func TestTodoGetByID(t *testing.T) {
	todo, err := todoServ.GetByID(1)
	if err != nil {
		t.Fatalf("TodoGetByID has failed: %v", err)
	}

	assert.Equal(t, todoDTO.Body, todo.Body)
}

func TestTodoGetAll(t *testing.T) {
	todos, err := todoServ.GetAll()
	if err != nil {
		t.Fatalf("TodoGetAll has failed: %v", err)
	}

	assert.Greater(t, len(*todos), 0)
}

func TestTodoUpdateByID(t *testing.T) {
	todoDTO.ID = 1
	todoDTO.Body = "update success"
	todo, err := todoServ.UpdateByID(&todoDTO)
	if err != nil {
		t.Fatalf("TodoUpdateByID has failed: %v", err)
	}
	assert.Equal(t, todoDTO.ID, todo.ID)
	assert.Equal(t, todoDTO.Body, todo.Body)
}
