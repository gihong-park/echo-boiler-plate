package middlewares

import (
	"net/http"

	"blog_api/app/auth"
	"blog_api/app/auth/role"

	"github.com/labstack/echo/v4"
)

func TokenAuthMiddleware(r role.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := auth.TokenValidByRole(r)(c.Request())
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")
			}
			return next(c)
		}
	}
}
