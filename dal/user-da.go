package dal

import (
	"database/sql"
	"github.com/prattnj/fms-go/model"
)

// void insert(User), boolean validate(username, password), User find(username), String getGender(username), void clear

func U_insert(db *sql.DB, user model.User) {
	// todo insert user
}

func U_find(db *sql.DB, username string) model.User {
	// todo find user
	return model.User{}
}

func U_validate(db *sql.DB, username string, password string) bool {
	// todo validate
	return true
}

func U_getGender(db *sql.DB, username string) string {
	return "m"
}

func U_clear(db *sql.DB) error {
	stmt, err := db.Prepare("DELETE FROM user")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
