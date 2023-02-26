package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/prattnj/fms-go/model"
	"github.com/prattnj/fms-go/service"
)

func HandleLogin(c echo.Context) error {

	// Convert request body
	var req model.LoginRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(400, model.GenericResponse{Success: false, Message: "Error: improperly formatted request. Details: " + err.Error()})
	}

	// Perform login and return appropriate response
	resp := service.Login(req.Username, req.Password)
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}

func HandleRegister(c echo.Context) error {

	// Convert request body
	var req model.RegisterRequest
	err := json.NewDecoder(c.Request().Body).Decode(&req)
	if err != nil {
		return c.JSON(400, model.GenericResponse{Success: false, Message: "Error: improperly formatted request. Details: " + err.Error()})
	}

	// Perform register and return appropriate response
	resp := service.Register(req.Username, req.Password, req.Email, req.FirstName, req.LastName, req.Gender)
	if resp.Success {
		return c.JSON(200, resp)
	} else {
		if resp.Message == "Internal server error" {
			return c.JSON(500, resp)
		}
		return c.JSON(400, resp)
	}
}
