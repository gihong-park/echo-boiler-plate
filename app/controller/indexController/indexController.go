package indexController

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (controller *IndexController) IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
