package service

import "github.com/prattnj/fms-go/model"

func Fill(username string, generations string) model.GenericResponse {
	// todo error checking (including asserting username and generations are not null)
	// todo fill
	return model.GenericResponse{Success: true, Message: "Successfully added " + username + " and " + generations + " generations"}
}
