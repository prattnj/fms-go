package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/prattnj/fms-go/service"
)

func HandleClear(c echo.Context) error {

	// Perform clear and return appropriate response
	resp := service.Clear()
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}
