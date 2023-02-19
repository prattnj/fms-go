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

func E_clear(db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM event")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func E_clearForUser(db *sql.DB, username string) {
	// todo clear events for user
}
