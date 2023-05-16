//go:build wireinject
// +build wireinject

package authController

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitAuthController(db *gorm.DB) AuthController {
	wire.Build(NewAuthController, NewAuthService)

	return AuthController{}
}
