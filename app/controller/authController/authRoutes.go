package authController

import "github.com/labstack/echo/v4"

func (authController *AuthController) Routes(g *echo.Group) {
	//api/v1/auth
	g.POST("signUp", authController.SignUp)
	g.POST("signIn", authController.SignIn)
}
