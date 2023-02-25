package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/prattnj/fms-go/service"
)

func HandleEvent(c echo.Context) error {

	// Perform event service and return appropriate response
	resp := service.Event(c.Request().Header.Get("Authorization"))
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Error: bad token" {
			return c.JSON(400, resp)
		}
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}

func HandleEventID(c echo.Context) error {

	// Perform eventID service and return appropriate response
	resp := service.EventID(c.Request().Header.Get("Authorization"), c.Param("eventID"))
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Error: bad token" {
			return c.JSON(400, resp)
		}
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}
