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

func init() {
	database = db.GetDB("sqlite")
	database.Migrator().DropTable(&model.Todo{})
	database.AutoMigrate(&model.Todo{})
}

func TestTodoSave(t *testing.T) {
	var todoRepo TodoRepository = NewTodoRepository[*TodoRepositoryImpl]()
	todoRepo.SetDB(database)
	todoItem := dto.TodoDTO{Body: "save item"}
	result := todoRepo.Save(&todoItem)
	t.Log(result)
	assert.Equal(t, todoItem.Body, result.Body)
}

func TestTest(t *testing.T) {
	assert.Equal(t, "Hello", "Hello")
}
