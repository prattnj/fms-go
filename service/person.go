package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
)

func Person(authtoken string) model.PersonResponse {

	if authtoken == "" {
		return model.PersonResponse{Success: false, Message: tokenErrorStr}
	}

	db := dal.Db()
	if db == nil {
		return model.PersonResponse{Success: false, Message: serverErrorStr}
	}
	tx, err := db.Begin()
	if err != nil {
		return model.PersonResponse{Success: false, Message: serverErrorStr}
	}

	username, err := dal.T_getUsername(tx, authtoken)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.PersonResponse{Success: false, Message: serverErrorStr}
		}
		return model.PersonResponse{Success: false, Message: serverErrorStr}
	}
	if username == "" {
		return model.PersonResponse{Success: false, Message: tokenErrorStr}
	}
	persons, err := dal.P_getForUsername(tx, username)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.PersonResponse{Success: false, Message: serverErrorStr}
		}
		return model.PersonResponse{Success: false, Message: serverErrorStr}
	}

	if commitAndClose(tx, db, true) != nil {
		return model.PersonResponse{Success: false, Message: serverErrorStr}
	}

	return model.PersonResponse{
		Data:    persons,
		Success: true,
	}
}

func PersonID(authtoken string, personID string) model.PersonIDResponse {

	if authtoken == "" {
		return model.PersonIDResponse{Success: false, Message: tokenErrorStr}
	}
	if personID == "" {
		return model.PersonIDResponse{Success: false, Message: "Error: missing personID"}
	}

	db := dal.Db()
	if db == nil {
		return model.PersonIDResponse{Success: false, Message: serverErrorStr}
	}
	tx, err := db.Begin()
	if err != nil {
		return model.PersonIDResponse{Success: false, Message: serverErrorStr}
	}

	username, err := dal.T_getUsername(tx, authtoken)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.PersonIDResponse{Success: false, Message: serverErrorStr}
		}
		return model.PersonIDResponse{Success: false, Message: serverErrorStr}
	}
	if username == "" {
		return model.PersonIDResponse{Success: false, Message: tokenErrorStr}
	}
	person, err := dal.P_find(tx, personID)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.PersonIDResponse{Success: false, Message: serverErrorStr}
		}
		return model.PersonIDResponse{Success: false, Message: serverErrorStr}
	}
	if person.PersonID == "" {
		if commitAndClose(tx, db, false) != nil {
			return model.PersonIDResponse{Success: false, Message: serverErrorStr}
		}
		return model.PersonIDResponse{Success: false, Message: "Error: person does not exist"}
	}
	if person.AssociatedUsername != username {
		if commitAndClose(tx, db, false) != nil {
			return model.PersonIDResponse{Success: false, Message: serverErrorStr}
		}
		return model.PersonIDResponse{Success: false, Message: "Error: person does not belong to this user"}
	}

	if commitAndClose(tx, db, true) != nil {
		return model.PersonIDResponse{Success: false, Message: serverErrorStr}
	}

	return model.PersonIDResponse{
		AssociatedUsername: person.AssociatedUsername,
		PersonID:           person.PersonID,
		FirstName:          person.FirstName,
		LastName:           person.LastName,
		Gender:             person.Gender,
		FatherID:           person.FatherID,
		MotherID:           person.MotherID,
		SpouseID:           person.SpouseID,
		Success:            true,
	}
}
