package service

import "github.com/prattnj/fms-go/model"

func Clear() model.GenericResponse {
	// todo clear
	return model.GenericResponse{Success: true, Message: "Clear succeeded."}
}
