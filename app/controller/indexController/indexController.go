package indexController

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func (controller *IndexController) IndexHandler(c echo.Context) error {
	log.Infof("Hello, World to %s", c.RealIP())
	return c.String(http.StatusOK, "Hello, World!")
}
