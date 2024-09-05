package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gokuls-codes/go-shadcn/templates/pages"
	"github.com/gokuls-codes/go-shadcn/utils"
	"github.com/labstack/echo/v4"
)


func theme(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("theme")
		if err != nil {
			c.Set("theme", "dark")
			cookie := new(http.Cookie)
			cookie.Name = "theme"
			cookie.Value = "dark"
			cookie.Expires = time.Now().Add(24 * time.Hour)
			cookie.Path = "/"
			c.SetCookie(cookie)
			return next(c)
		}
		c.Set("theme", cookie.Value)
		return next(c)
	}
}

func main() {
	app := echo.New()

	app.Use(theme)

	app.GET("/", func(c echo.Context) error {
		theme := c.Get("theme").(string)
		return utils.Render(c, pages.HomePage(theme == "dark"))
	})

	log.Fatal(app.Start(":8080"))
}