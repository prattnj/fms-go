package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
	"github.com/prattnj/fms-go/service"
)

func HandleLoad(c echo.Context) error {

	if c.Request().Header.Get("Authorization") != dal.GetPassword() {
		service.Log(c.Path(), c.RealIP(), false)
		return c.JSON(401, model.GenericResponse{Success: false, Message: "Bad token"})
	}

	// Convert request body
	var req model.LoadRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		service.Log(c.Path(), c.RealIP(), false)
		return c.JSON(400, model.GenericResponse{Success: false, Message: "Error: improperly formatted request. Details: " + err.Error()})
	}

	// Perform load and return appropriate response
	resp := service.Load(req.Users, req.Persons, req.Events)
	service.Log(c.Path(), c.RealIP(), resp.Success)
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}
