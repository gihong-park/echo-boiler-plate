package todoRepository

import (
	"testing"

	"blog_api/app/controller/dto"
	"blog_api/app/db"

	"blog_api/app/model"

	_ "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var database *gorm.DB
var todoItem dto.TodoDTO = dto.TodoDTO{Body: "save item"}
var todoRepo TodoRepository

func init() {
	database = db.GetDB("sqlite")
	database.Migrator().DropTable(&model.Todo{})
	database.AutoMigrate(&model.Todo{})
	todoRepo = &TodoRepositoryImpl{database}
	todoRepo.Save(&todoItem)
}

func TestTodoSave(t *testing.T) {
	result, err := todoRepo.Save(&todoItem)
	if err != nil {
		t.Fatalf("TodoSave has failed: %v", err)
	}
	t.Log(result)
	assert.Equal(t, todoItem.Body, result.Body)
}

func TestTodoGetByID(t *testing.T) {
	id := uint(1)
	result, err := todoRepo.GetByID(id)
	if err != nil {
		t.Fatalf("TodoGetByID has failed: %v", err)
	}
	t.Log(result)
	assert.Equal(t, todoItem.Body, result.Body)
	assert.Equal(t, id, result.ID)
}

func TestTodoGetAll(t *testing.T) {
	todos, err := todoRepo.GetAll()
	if err != nil {
		t.Fatalf("TodoGetAll has failed: %v", err)
	}

	t.Log(todos)
	assert.Greater(t, len(*todos), 0)
}

func TestTodoUpdateByID(t *testing.T) {
	todoItem.Body = "updated Item"
	todoItem.ID = 1
	todo, err := todoRepo.UpdateByID(&todoItem)
	if err != nil {
		t.Fail()
		t.Fatalf("TodoUpdateByID: %v", err)
	}
	if assert.NotEmpty(t, todo) {
		t.Log(todo)
		assert.Equal(t, todoItem.Body, todo.Body)
		assert.Equal(t, todoItem.ToModel().ID, todo.ID)
	}
}

func TestSetDB(t *testing.T) {
	todoRepo.SetDB(database)

	assert.Equal(t, database, todoRepo.GetDB())
}

func TestGetDB(t *testing.T) {
	assert.Equal(t, database, todoRepo.GetDB())
}
