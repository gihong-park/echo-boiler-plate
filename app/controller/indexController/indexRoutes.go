package indexController

import "github.com/labstack/echo/v4"

func (indexController IndexController) Routes(g *echo.Group) (routes []*echo.Route) {
	routes = append(routes, g.GET("", indexController.indexHandler))

	return routes
}
