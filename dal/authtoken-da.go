package dal

import (
	"database/sql"
	"fmt"
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
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var authToken model.AuthToken
	for rows.Next() {
		err = rows.Scan(&authToken.AuthToken, &authToken.Username)
		if err != nil {
			return model.AuthToken{}, err
		}
		return authToken, nil
	}
	return model.AuthToken{}, nil
}

func T_getUsername(tx *sql.Tx, token string) (string, error) {
	authToken, err := T_find(tx, token)
	if err != nil {
		return "", err
	}
	return authToken.Username, nil
}

func T_clear(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM authtoken;")
	if err != nil {
		return err
	}
	return nil
}
