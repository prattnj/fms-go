package dal

import (
	"database/sql"
	"fmt"
	"github.com/prattnj/fms-go/model"
)

func P_insert(tx *sql.Tx, person model.Person) error {
	_, err := tx.Exec("INSERT INTO person (personID, associatedUsername, firstName, lastName, gender, fatherID, motherID, spouseID) VALUES(?, ?, ?, ?, ?, ?, ?, ?);",
		person.PersonID, person.AssociatedUsername, person.FirstName, person.LastName, person.Gender, person.FatherID, person.MotherID, person.SpouseID)
	if err != nil {
		return err
	}
	return nil
}
func P_find(tx *sql.Tx, id string) (model.Person, error) {
	rows, err := tx.Query("SELECT * FROM person WHERE personID = ?;", id)
	if err != nil {
		return model.Person{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var person model.Person
	for rows.Next() {
		err = rows.Scan(&person.PersonID, &person.AssociatedUsername, &person.FirstName, &person.LastName, &person.Gender, &person.FatherID, &person.MotherID, &person.SpouseID)
		if err != nil {
			return model.Person{}, err
		}
		return person, nil
	}
	return model.Person{}, nil
}

func P_getForUsername(tx *sql.Tx, username string) ([]model.Person, error) {
	rows, err := tx.Query("SELECT * FROM person WHERE associatedUsername = ?;", username)
	if err != nil {
		return []model.Person{}, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Error closing rows:", err)
		}
	}(rows)
	var people []model.Person
	for rows.Next() {
		var person model.Person
		err = rows.Scan(&person.PersonID, &person.AssociatedUsername, &person.FirstName, &person.LastName, &person.Gender, &person.FatherID, &person.MotherID, &person.SpouseID)
		if err != nil {
			return []model.Person{}, err
		}
		people = append(people, person)
	}
	return people, nil
}

func P_clear(tx *sql.Tx) error {
	_, err := tx.Exec("DELETE FROM person;")
	if err != nil {
		return err
	}
	return nil
}

func P_clearForUser(tx *sql.Tx, username string) error {
	_, err := tx.Exec("DELETE FROM person WHERE associatedUsername = ?;", username)
	if err != nil {
		return err
	}
	return nil
}
