package app

import (
	"blog_api/app/controller"
	"blog_api/app/db"
	"blog_api/app/model"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type App struct {
	db     *gorm.DB
	server *echo.Echo
}

func init() {
	err := godotenv.Load(".env")
	fmt.Println(err)
}

func Init() {
	echoServer := echo.New()
	app := App{db: db.GetDB("sqlite"), server: echoServer}

	cont := controller.InitController(app.db)

	cont.Routes(app.server)

	app.db.Debug().AutoMigrate(&model.Todo{})
	app.server.Use(middleware.CORS())

	app.server.Logger.Fatal(app.server.Start(":8000"))
}
