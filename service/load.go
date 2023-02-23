package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
)

func Load(users []model.User, persons []model.Person, events []model.Event) model.GenericResponse {

	// Maximum number of objects allowed in request body
	var maxObj = 1000
	if len(users)+len(persons)+len(events) > maxObj {
		return model.GenericResponse{Success: false, Message: "Request body too large"}
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
			err := tx.Rollback()
			if err != nil {
				return serverError
			}
			return serverError
		}
	}
	for person := range persons {
		err := dal.P_insert(tx, persons[person])
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return serverError
			}
			return serverError
		}
	}
	for event := range events {
		err := dal.E_insert(tx, events[event])
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return serverError
			}
			return serverError
		}
	}

	err = tx.Commit()
	if err != nil {
		return serverError
	}
	err = dal.DbClose(db)
	if err != nil {
		return serverError
	}
	return model.GenericResponse{Success: true, Message: "Successfully added " + string(rune(len(users))) + " users, " +
		string(rune(len(persons))) + " persons, and " + string(rune(len(events))) + " events to the database."}
}
