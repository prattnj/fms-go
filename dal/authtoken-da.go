package dal

import (
	"database/sql"
	"github.com/prattnj/fms-go/model"
)

func T_insert(tx *sql.Tx, token model.AuthToken) error {
	_, err := tx.Exec("INSERT INTO authtoken (authtoken, username) VALUES(?,?);", token.AuthToken, token.Username)
	if err != nil {
		return err
	}
	return nil
}

func T_find(tx *sql.Tx, token string) (model.AuthToken, error) {
	rows, err := tx.Query("SELECT * FROM authtoken WHERE authtoken = ?;", token)
	if err != nil {
		return model.AuthToken{}, err
	}
	var authToken string
	var username string
	for rows.Next() {
		err = rows.Scan(&authToken, &username)
		if err != nil {
			return model.AuthToken{}, err
		}
		return model.AuthToken{AuthToken: authToken, Username: username}, nil
	}
	return model.AuthToken{}, nil
}

func T_clear(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM authtoken;")
	if err != nil {
		return err
	}
	return nil
}
