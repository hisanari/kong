package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Name string `json:"name"`
}

type Users []User

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, World!!")
	})

	e.GET("/users", func(c echo.Context) error {
		users := Users{
			User{
				Name: "bob",
			},
			User{
				Name: "Alice",
			},
		}
		return c.JSON(http.StatusOK, users)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
