package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/prattnj/fms-go/service"
)

func HandlePerson(c echo.Context) error {

	// Perform event service and return appropriate response
	resp := service.Person(c.Request().Header.Get("Authorization"))
	service.Log(c.Path(), c.RealIP(), resp.Success)
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Error: bad token" {
			return c.JSON(401, resp)
		}
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}

func HandlePersonID(c echo.Context) error {

	// Perform eventID service and return appropriate response
	resp := service.PersonID(c.Request().Header.Get("Authorization"), c.Param("personID"))
	service.Log(c.Path(), c.RealIP(), resp.Success)
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Error: bad token" {
			return c.JSON(401, resp)
		}
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}
