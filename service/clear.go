package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
)

func Clear() model.GenericResponse {

	var serverError = model.GenericResponse{Success: false, Message: "Internal server error"}

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
		err := tx.Rollback()
		if err != nil {
			return serverError
		}
		return serverError
	}

	err = tx.Commit()
	if err != nil {
		return serverError
	}
	err = dal.DbClose(db)
	if err != nil {
		return serverError
	}
	return model.GenericResponse{Success: true, Message: "Clear succeeded"}
}
