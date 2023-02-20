package dal

import (
	"database/sql"
	"github.com/prattnj/fms-go/model"
)

func E_insert(db *sql.DB, event model.Event) {
	// todo insert event
}

func E_find(db *sql.DB, eventID string) model.Event {
	// todo find event
	return model.Event{}
}

func E_findBirth(db *sql.DB, personID string) model.Event {
	// todo find birth event
	return model.Event{}
}

func E_getForUsername(db *sql.DB, username string) []model.Event {
	// todo get events for username
	return []model.Event{}
}

func E_clear(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM event;")
	if err != nil {
		return err
	}
	return nil
}

func E_clearForUser(tx *sql.Tx, username string) error {
	_, err := tx.Exec("DELETE FROM event WHERE username = ?;", username)
	if err != nil {
		return err
	}
	return nil
}
