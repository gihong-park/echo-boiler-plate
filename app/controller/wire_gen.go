// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package controller

import (
	"gorm.io/gorm"
)

// Injectors from wire.go:

func InitController(db *gorm.DB) Controller {
	controller := NewController(db)
	return controller
}
