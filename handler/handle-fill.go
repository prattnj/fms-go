package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/prattnj/fms-go/service"
)

func HandleFill(c echo.Context) error {

	// Perform fill and return appropriate response
	username := c.Param("username")
	generations := c.Param("generations")
	if generations == "" {
		generations = "4"
	}
	resp := service.Fill(username, generations)
	service.Log(compileFillPath(username, generations), c.RealIP(), resp.Success)
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}

func compileFillPath(username string, generations string) string {
	return "/fill/" + username + "/" + generations
}
