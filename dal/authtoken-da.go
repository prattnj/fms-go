package dal

import (
	"database/sql"
	"github.com/prattnj/fms-go/model"
)

func T_insert(db *sql.DB, token model.AuthToken) {
	// todo add token
}

func T_find(db *sql.DB, token string) model.AuthToken {
	// todo find token
	return model.AuthToken{}
}

func T_getUsername(db *sql.DB, token string) string {
	// todo get username
	return ""
}

func T_clear(db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM authtoken")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
