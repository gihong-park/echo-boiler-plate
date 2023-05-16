package util

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("[Fatal] fail to get root directory: %w", err)
	}
	return dir
}

func IsTest() bool {
	return strings.HasSuffix(os.Args[0], ".test")
}

func NewServer() *echo.Echo {
	echoServer := echo.New()

	echoServer.JSONSerializer = JSONIterSerializer{}

	return echoServer
}
