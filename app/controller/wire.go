//go:build wireinject
// +build wireinject

package controller

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitController(db *gorm.DB) Controller {
	wire.Build(NewController)

	return Controller{}
}
