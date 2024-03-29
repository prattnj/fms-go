package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
	"strconv"
)

func Load(users []model.User, persons []model.Person, events []model.Event) model.GenericResponse {

	// Maximum number of objects allowed in request body
	var maxObj = 1000
	if len(users)+len(persons)+len(events) > maxObj {
		return model.GenericResponse{Success: false, Message: "Error: request body too large"}
	}

	// Clear the database
	clearResp := Clear()
	if clearResp.Success == false {
		return clearResp
	}

	// Load the database
	db := dal.Db()
	if db == nil {
		return serverError
	}
	tx, err := db.Begin()
	if err != nil {
		return serverError
	}
	for user := range users {
		err := dal.U_insert(tx, users[user])
		if err != nil {
			if commitAndClose(tx, db, false) != nil {
				return serverError
			}
			return serverError
		}
	}
	for person := range persons {
		err := dal.P_insert(tx, &persons[person])
		if err != nil {
			if commitAndClose(tx, db, false) != nil {
				return serverError
			}
			return serverError
		}
	}
	for event := range events {
		err := dal.E_insert(tx, &events[event])
		if err != nil {
			if commitAndClose(tx, db, false) != nil {
				return serverError
			}
			return serverError
		}
	}

	if commitAndClose(tx, db, true) != nil {
		return serverError
	}

	return model.GenericResponse{Success: true, Message: "Successfully added " + strconv.Itoa(len(users)) + " users, " +
		strconv.Itoa(len(persons)) + " persons, and " + strconv.Itoa(len(events)) + " events to the database."}
}
