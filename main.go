package main

import (
	"github.com/labstack/echo/v4"
	"github.com/prattnj/fms-go/handler"
	"log"
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

	err := e.StartAutoTLS(":8443")
	if err != nil {
		log.Fatal(err)
	}
}
