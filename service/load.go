package service

import (
	"github.com/prattnj/fms-go/model"
)

func Load(users []model.User, persons []model.Person, events []model.Event) model.GenericResponse {
	// todo load
	return model.GenericResponse{Success: true, Message: "Load successful"}
}
