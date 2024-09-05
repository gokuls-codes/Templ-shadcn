package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	log.Fatal(app.Start(":8080"))
}