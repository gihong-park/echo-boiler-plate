package util

import (
	"log"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Println("[Fatal] fail to get root directory")
		log.Fatal(err)
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
