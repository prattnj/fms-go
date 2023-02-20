package dal

import (
	"database/sql"
	"github.com/prattnj/fms-go/model"
)

func P_insert(db *sql.DB, person model.Person) {
	// todo add person
}
func P_find(db *sql.DB, personID string) model.Person {
	// todo find person
	return model.Person{}
}

func P_getForUsername(db *sql.DB, username string) []model.Person {
	// todo get person for username
	return []model.Person{}
}

func P_clear(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM person;")
	if err != nil {
		return err
	}
	return nil
}

func P_clearForUser(db *sql.DB, username string) {
	// todo clear person table for user
}
