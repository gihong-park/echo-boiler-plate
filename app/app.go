package app

import (
	"blog_api/app/controller"
	"blog_api/app/db"
	"blog_api/app/model"
	"blog_api/app/util"
	"fmt"
	"os"
	"path/filepath"

	promMW "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type App struct {
	db     *gorm.DB
	server *echo.Echo
}

func Init() {
	echoServer := util.NewServer()
	app := App{db: db.GetDB("sqlite"), server: echoServer}
	DefaultRoutes(app)

	f := RegistMiddlewares(app.server)
	defer f.Close()

	app.db.Debug().AutoMigrate(&model.Todo{})

	p := promMW.NewPrometheus("blog_api", nil)
	p.Use(app.server)
	app.server.HTTPErrorHandler = echoServer.DefaultHTTPErrorHandler

	app.server.Logger.Fatal(app.server.Start(":" + os.Getenv("API_PORT")))
}

func LoggerConfigure(config middleware.LoggerConfig) (middleware.LoggerConfig, *os.File) {
	absPath, _ := filepath.Abs(".")
	fmt.Println("root folder:", absPath)
	f, _ := os.Create(absPath + "/blog_api.log")

	config.Output = f
	return config, f
}

func DefaultRoutes(app App) {
	cont := controller.InitController(app.db)
	cont.Routes(app.server)
}

func RegistMiddlewares(e *echo.Echo) *os.File {
	config, f := LoggerConfigure(middleware.DefaultLoggerConfig)

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(config))
	e.Use(middleware.Recover())
	// e.Use(middleware.CSRF())

	return f
}
