package dal

import (
	"database/sql"
	"fmt"
	"github.com/prattnj/fms-go/model"
)

func U_insert(tx *sql.Tx, user model.User) error {
	_, err := tx.Exec("INSERT INTO user (username, password, email, firstName, lastName, gender, personID) VALUES(?, ?, ?, ?, ?, ?, ?);",
		user.Username, user.Password, user.Email, user.FirstName, user.LastName, user.Gender, user.PersonID)
	if err != nil {
		return err
	}
	return nil
}

func U_find(tx *sql.Tx, username string) (model.User, error) {
	rows, err := tx.Query("SELECT * FROM user WHERE username = ?;", username)
	if err != nil {
		return model.User{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Gender, &user.PersonID)
		if err != nil {
			return model.User{}, err
		}
		return user, nil
	}
	return model.User{}, nil
}

func U_validate(tx *sql.Tx, username string, password string) (bool, error) {
	rows, err := tx.Query("SELECT * FROM user WHERE username = ? AND password = ?;", username, password)
	if err != nil {
		return false, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	return rows.Next(), nil
}

func U_getGender(tx *sql.Tx, username string) (string, error) {
	user, err := U_find(tx, username)
	if err != nil {
		return "", err
	}
	return user.Gender, nil
}

func U_getCount(tx *sql.Tx) (int, error) {
	rows, err := tx.Query("SELECT COUNT(*) FROM user;")
	if err != nil {
		return 0, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
		return count, nil
	}
	return 0, nil
}

func U_clear(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM user;")
	if err != nil {
		return err
	}
	return nil
}
