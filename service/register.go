package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
)

func Register(username string, password string, email string, firstname string, lastname string, gender string) model.LoginResponse {

	if gender != "m" && gender != "f" {
		return model.LoginResponse{Success: false, Message: "Invalid gender (must be 'm' or 'f')"}
	}

	db := dal.Db()
	if db == nil {
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}
	tx, err := db.Begin()
	if err != nil {
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}
	user, err := dal.U_find(tx, username)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return model.LoginResponse{Success: false, Message: serverErrorStr}
		}
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}
	if user.Username != "" {
		return model.LoginResponse{Success: false, Message: "Username already exists"}
	}

	newUser := model.User{
		Username:  username,
		Password:  password,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Gender:    gender,
		PersonID:  generateID(defaultIDLength),
	}
	err = dal.U_insert(tx, newUser)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return model.LoginResponse{Success: false, Message: serverErrorStr}
		}
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}

	err = generateData(4, gender)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return model.LoginResponse{Success: false, Message: serverErrorStr}
		}
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}

	for person := range fillPeople {
		err := dal.P_insert(tx, fillPeople[person])
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return model.LoginResponse{Success: false, Message: serverErrorStr}
			}
			return model.LoginResponse{Success: false, Message: serverErrorStr}
		}
	}
	for event := range fillEvents {
		err := dal.E_insert(tx, fillEvents[event])
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return model.LoginResponse{Success: false, Message: serverErrorStr}
			}
			return model.LoginResponse{Success: false, Message: serverErrorStr}
		}
	}

	token := "auth-" + generateID(defaultIDLength-5)
	err = dal.T_insert(tx, model.AuthToken{
		AuthToken: token,
		Username:  username,
	})
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return model.LoginResponse{Success: false, Message: serverErrorStr}
		}
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}

	err = tx.Commit()
	if err != nil {
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}
	err = dal.DbClose(db)
	if err != nil {
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}

	return model.LoginResponse{
		AuthToken: token,
		UserName:  username,
		PersonID:  newUser.PersonID,
		Success:   true,
		Message:   "",
	}
}
