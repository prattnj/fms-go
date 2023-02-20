package dal

import (
	"database/sql"
	"github.com/prattnj/fms-go/model"
)

func T_insert(db *sql.DB, token model.AuthToken) error {
	// todo add token
	stmt, err := db.Prepare("INSERT INTO Authtoken (authtoken, username) VALUES(?,?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(token.AuthToken, token.Username)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func T_find(db *sql.DB, token string) (model.AuthToken, error) {
	// todo find token
	stmt, err := db.Prepare("SELECT * FROM Authtoken WHERE authtoken = ?;")
	if err != nil {
		return model.AuthToken{}, err
	}
	var authToken string
	var username string
	err = stmt.QueryRow(token).Scan(&authToken, &username)
	if err != nil {
		return model.AuthToken{}, err
	}
	return model.AuthToken{AuthToken: authToken, Username: username}, nil
}

func T_clear(db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM authtoken;")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
