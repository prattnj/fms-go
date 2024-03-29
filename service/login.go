package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
)

func Login(username string, password string) model.LoginResponse {

	if username == "" || password == "" {
		return model.LoginResponse{Success: false, Message: "Error: missing username or password"}
	}

	// Validate username and password
	db := dal.Db()
	if db == nil {
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}
	tx, err := db.Begin()
	if err != nil {
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}
	valid, err := dal.U_validate(tx, username, password)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.LoginResponse{Success: false, Message: serverErrorStr}
		}
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}
	if !valid {
		return model.LoginResponse{Success: false, Message: "Error: invalid username or password"}
	}

	// Generate and return auth token
	user, err := dal.U_find(tx, username)
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.LoginResponse{Success: false, Message: serverErrorStr}
		}
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}
	token := "auth-" + generateID(defaultIDLength-5)
	err = dal.T_insert(tx, model.AuthToken{
		AuthToken: token,
		Username:  username,
	})
	if err != nil {
		if commitAndClose(tx, db, false) != nil {
			return model.LoginResponse{Success: false, Message: serverErrorStr}
		}
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}
	if commitAndClose(tx, db, true) != nil {
		return model.LoginResponse{Success: false, Message: serverErrorStr}
	}

	return model.LoginResponse{
		AuthToken: token,
		UserName:  username,
		PersonID:  user.PersonID,
		Success:   true,
		Message:   "",
	}
}
