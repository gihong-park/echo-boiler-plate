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
)

func init() {
	database = db.GetDB("sqlite")
	database.Migrator().DropTable(&model.Todo{})
	database.AutoMigrate(&model.Todo{})
}

func TestTodoService(t *testing.T) {
	todoServ = NewTodoService[*TodoServiceImpl]()
	var todoRepo todoRepository.TodoRepository = &todoRepository.TodoRepositoryImpl{DB: database}
	todoServ.SetRepo(todoRepo)
	todoDTO := dto.TodoDTO{Body: "save item"}
	todoModel, err := todoServ.Save(&todoDTO)
	if err != nil {
		t.Fatalf("TodoService has failed: %v", err)
	}

	assert.Equal(t, todoRepo, todoServ.SetRepo(todoRepo))
	t.Log(todoModel)
	assert.Equal(t, todoDTO.Body, todoModel.Body)
}
