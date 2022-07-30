package indexController

import "github.com/labstack/echo/v4"

type IndexController struct {
}

func (indexController IndexController) Routes(g *echo.Group) (routes []*echo.Route) {
	routes = append(routes, g.GET("", indexController.IndexHandler))

	return routes
}
