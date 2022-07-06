package todoController

import (
	"blog_api/app/controller/dto"
	"blog_api/app/db"
	"blog_api/app/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestTodoController(t *testing.T) {
	e := echo.New()
	todoCont := InitTodoController(db.GetDB("sqlite"))

	bodyContent := "save test1"
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
