package indexController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexController struct {
}

func (controller *IndexController) indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
