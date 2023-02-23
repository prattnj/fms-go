package dal

import (
	"database/sql"
	"fmt"
	"github.com/prattnj/fms-go/model"
)

func E_insert(tx *sql.Tx, event *model.Event) error {
	_, err := tx.Exec("INSERT INTO event (eventID, associatedUsername, personID, latitude, longitude, country, city, eventType, year) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);",
		event.EventID, event.AssociatedUsername, event.PersonID, event.Latitude, event.Longitude, event.Country, event.City, event.EventType, event.Year)
	if err != nil {
		return err
	}
	return nil
}

func E_find(tx *sql.Tx, eventID string) (model.Event, error) {
	rows, err := tx.Query("SELECT * FROM event WHERE eventID = ?;", eventID)
	if err != nil {
		return model.Event{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var event model.Event
	for rows.Next() {
		err = rows.Scan(&event.EventID, &event.AssociatedUsername, &event.PersonID, &event.Latitude, &event.Longitude, &event.Country, &event.City, &event.EventType, &event.Year)
		if err != nil {
			return model.Event{}, err
		}
		return event, nil
	}
	return model.Event{}, nil
}

func E_findBirth(tx *sql.Tx, personID string) (model.Event, error) {
	rows, err := tx.Query("SELECT * FROM event WHERE personID = ? AND (eventType = ? OR eventType = ? OR eventType = ?);", personID, "birth", "BIRTH", "Birth")
	if err != nil {
		return model.Event{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var event model.Event
	for rows.Next() {
		err = rows.Scan(&event.EventID, &event.AssociatedUsername, &event.PersonID, &event.Latitude, &event.Longitude, &event.Country, &event.City, &event.EventType, &event.Year)
		if err != nil {
			return model.Event{}, err
		}
		return event, nil
	}
	return model.Event{}, nil
}

func E_getForUsername(tx *sql.Tx, username string) ([]model.Event, error) {
	rows, err := tx.Query("SELECT * FROM event WHERE associatedUsername = ?;", username)
	if err != nil {
		return []model.Event{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var events []model.Event
	for rows.Next() {
		var event model.Event
		err = rows.Scan(&event.EventID, &event.AssociatedUsername, &event.PersonID, &event.Latitude, &event.Longitude, &event.Country, &event.City, &event.EventType, &event.Year)
		if err != nil {
			return []model.Event{}, err
		}
		events = append(events, event)
	}
	return events, nil
}

func E_clear(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM event;")
	if err != nil {
		return err
	}
	return nil
}

func E_clearForUser(tx *sql.Tx, username string) error {
	_, err := tx.Exec("DELETE FROM event WHERE associatedUsername = ?;", username)
	if err != nil {
		return err
	}
	return nil
}
