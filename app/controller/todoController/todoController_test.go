package todoController

import (
	"blog_api/app/controller/dto"
	"blog_api/app/db"
	"blog_api/app/model"
	"blog_api/app/util"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	database := db.GetDB("sqlite")
	database.Migrator().DropTable(&model.Todo{})
	database.AutoMigrate(&model.Todo{})
}

func TestTodoController(t *testing.T) {
	e := util.NewServer()
	todoCont := InitTodoController(db.GetDB("sqlite"))

	bodyContent := "save item"
	todoDTO := dto.TodoDTO{Body: bodyContent}
	todoJson, _ := json.Marshal(todoDTO)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(todoJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/todo")

	todoCont.SaveHandler(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	todoModel := new(model.Todo)
	json.Unmarshal(rec.Body.Bytes(), todoModel)
	assert.Equal(t, bodyContent, todoModel.Body)
}

type Map map[string]interface{}

func TestTodoSaveController(t *testing.T) {
	e := util.NewServer()
	todoCont := InitTodoController(db.GetDB("sqlite"))

	todoDTO := dto.TodoDTO{Body: ""}
	todoJson, _ := json.Marshal(todoDTO)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(todoJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/todo")

	err := todoCont.SaveHandler(c)
	if assert.Error(t, err) {
		if he, ok := err.(*echo.HTTPError); ok {
			assert.Equal(t, http.StatusBadRequest, he.Code)
			assert.ErrorContains(t, he, "code=")
		}
	}
}

func TestTest(t *testing.T) {
	assert.Equal(t, "hello", "hello")
}
