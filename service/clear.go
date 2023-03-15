package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
)

func Clear() model.GenericResponse {

	db := dal.Db()
	if db == nil {
		return serverError
	}
	tx, err := db.Begin()
	if err != nil {
		return serverError
	}

	err1 := dal.T_clear(tx)
	err2 := dal.U_clear(tx)
	err3 := dal.P_clear(tx)
	err4 := dal.E_clear(tx)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		if commitAndClose(tx, db, false) != nil {
			return serverError
		}
		return serverError
	}

	if commitAndClose(tx, db, true) != nil {
		return serverError
	}
	return model.GenericResponse{Success: true, Message: "Clear succeeded"}
}
