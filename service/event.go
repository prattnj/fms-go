package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
)

func Event(authtoken string) model.EventResponse {

	if authtoken == "" {
		return model.EventResponse{Success: false, Message: tokenErrorStr}
	}

	db := dal.Db()
	if db == nil {
		return model.EventResponse{Success: false, Message: serverErrorStr}
	}
	tx, err := db.Begin()
	if err != nil {
		return model.EventResponse{Success: false, Message: serverErrorStr}
	}

	username, err := dal.T_getUsername(tx, authtoken)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.EventResponse{Success: false, Message: serverErrorStr}
		}
		return model.EventResponse{Success: false, Message: serverErrorStr}
	}
	if username == "" {
		if commitAndClose(tx, db, false) != nil {
			return model.EventResponse{Success: false, Message: serverErrorStr}
		}
		return model.EventResponse{Success: false, Message: tokenErrorStr}
	}
	events, err := dal.E_getForUsername(tx, username)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.EventResponse{Success: false, Message: serverErrorStr}
		}
		return model.EventResponse{Success: false, Message: serverErrorStr}
	}

	if commitAndClose(tx, db, true) != nil {
		return model.EventResponse{Success: false, Message: serverErrorStr}
	}

	return model.EventResponse{
		Data:    events,
		Success: true,
	}
}

func EventID(authtoken string, eventID string) model.EventIDResponse {

	if authtoken == "" {
		return model.EventIDResponse{Success: false, Message: tokenErrorStr}
	}
	if eventID == "" {
		return model.EventIDResponse{Success: false, Message: "Error: missing eventID"}
	}

	db := dal.Db()
	if db == nil {
		return model.EventIDResponse{Success: false, Message: serverErrorStr}
	}
	tx, err := db.Begin()
	if err != nil {
		return model.EventIDResponse{Success: false, Message: serverErrorStr}
	}

	username, err := dal.T_getUsername(tx, authtoken)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.EventIDResponse{Success: false, Message: serverErrorStr}
		}
		return model.EventIDResponse{Success: false, Message: serverErrorStr}
	}
	if username == "" {
		return model.EventIDResponse{Success: false, Message: tokenErrorStr}
	}
	event, err := dal.E_find(tx, eventID)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.EventIDResponse{Success: false, Message: serverErrorStr}
		}
		return model.EventIDResponse{Success: false, Message: serverErrorStr}
	}
	if event.EventID == "" {
		if commitAndClose(tx, db, false) != nil {
			return model.EventIDResponse{Success: false, Message: serverErrorStr}
		}
		return model.EventIDResponse{Success: false, Message: "Error: event does not exist"}
	}
	if event.AssociatedUsername != username {
		if commitAndClose(tx, db, false) != nil {
			return model.EventIDResponse{Success: false, Message: serverErrorStr}
		}
		return model.EventIDResponse{Success: false, Message: "Error: event does not belong to this user"}
	}

	if commitAndClose(tx, db, true) != nil {
		return model.EventIDResponse{Success: false, Message: serverErrorStr}
	}

	return model.EventIDResponse{
		AssociatedUsername: event.AssociatedUsername,
		EventID:            event.EventID,
		PersonID:           event.PersonID,
		Latitude:           event.Latitude,
		Longitude:          event.Longitude,
		Country:            event.Country,
		City:               event.City,
		EventType:          event.EventType,
		Year:               event.Year,
		Success:            true,
	}
}
