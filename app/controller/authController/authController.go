package authController

import (
	"blog_api/app/auth"
	"blog_api/app/controller/dto"
	"blog_api/app/security"
	"blog_api/app/service/authService"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AuthController struct {
	AuthServ authService.AuthService
}

func (authController *AuthController) SignUp(c echo.Context) error {
	ur := new(dto.UserRequest)
	c.Bind(ur)
	user, err := authController.AuthServ.Save(ur)
	if err != nil {
		log.Errorf("fail to save user: %+v", err)
		return err
	}
	return c.JSON(http.StatusCreated, dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role.String(),
	})
}

func (authController *AuthController) SignIn(c echo.Context) error {
	signIn := new(dto.SignIn)
	c.Bind(signIn)
	user, err := authController.AuthServ.GetByEmail(signIn.Email)
	if err != nil {
		log.Info(err)
		return err
	}
	if err = security.VerifyPassword(user.Password, signIn.Password); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}

	token, err := auth.CreateToken(uint32(user.ID), user.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"email": user.Email,
		"name":  user.Name,
		"role":  user.Role.String(),
		"token": token,
	})
}
