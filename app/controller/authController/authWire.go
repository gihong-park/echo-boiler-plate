package authController

import (
	"blog_api/app/service/authService"

	"gorm.io/gorm"
)

func NewAuthController(s authService.AuthService) AuthController {
	return AuthController{s}
}

func NewAuthService(db *gorm.DB) authService.AuthService {
	todoServ := &authService.AuthServiceImpl{}
	todoServ.SetDB(db)
	return todoServ
}
