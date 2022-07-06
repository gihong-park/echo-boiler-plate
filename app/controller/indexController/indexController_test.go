package indexController

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	e := echo.New()
	var controller IndexController = IndexController{}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1")

	controller.indexHandler(c)
	t.Log(c.Path())

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, World!", rec.Body.String())
}

func TestIndexRoutes(t *testing.T) {
	e := echo.New()

	g := e.Group("/api/v1")

	indexCont := IndexController{}

	routes := indexCont.Routes(g)
	assert.Equal(t, e.Routes(), routes)
}
