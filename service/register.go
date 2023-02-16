package service

import "github.com/prattnj/fms-go/model"

func Register(username string, password string, email string, firstname string, lastname string, gender string) model.LoginResponse {
	// todo register
	return model.LoginResponse{Success: false, Message: "Not implemented"}
}
