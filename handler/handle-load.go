package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/prattnj/fms-go/model"
	"github.com/prattnj/fms-go/service"
)

func HandleLoad(c echo.Context) error {

	// Convert request body
	var req model.LoadRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(400, model.GenericResponse{Success: false, Message: err.Error()})
	}

	// Perform load and return appropriate response
	resp := service.Load(req.Users, req.Persons, req.Events)
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}
