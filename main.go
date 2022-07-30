package main

import (
	"blog_api/app"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[Fatal] .env file can't load")
	}
}

func main() {
	app.Init()
}
