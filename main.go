package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prattnj/fms-go/handler"
	"log"
	"net/http"
)

func main() {
	e := echo.New()

	e.HTTPErrorHandler = handler.HandleNotFound
	e.Static("/", "./web")
	e.POST("/clear", handler.HandleClear)
	e.GET("/event", handler.HandleEvent)
	e.GET("/event/:eventID", handler.HandleEventID)
	e.POST("/fill/:username/:generations", handler.HandleFill)
	e.POST("/load", handler.HandleLoad)
	e.GET("/person", handler.HandlePerson)
	e.GET("/person/:personID", handler.HandlePersonID)
	e.POST("/user/login", handler.HandleLogin)
	e.POST("/user/register", handler.HandleRegister)
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `
			<h1>Welcome to Echo!</h1>
			<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
		`)
	})

	err := e.StartAutoTLS(":3003")
	if err != nil {
		log.Fatal(err)
	}
}
