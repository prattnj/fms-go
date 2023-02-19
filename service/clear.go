package service

import (
	"fmt"
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
)

func Clear() model.GenericResponse {

	db := dal.Db()
	if db == nil {
		return model.GenericResponse{Success: false, Message: "Internal server error"}
	}

	fmt.Println("Clearing tables...")
	err1 := dal.T_clear(db)
	err2 := dal.U_clear(db)
	err3 := dal.P_clear(db)
	err4 := dal.E_clear(db)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		return model.GenericResponse{Success: false, Message: "Internal server error"}
	}

	err := dal.DbClose(db)
	if err != nil {
		return model.GenericResponse{Success: false, Message: "Internal server error"}
	}
	return model.GenericResponse{Success: true, Message: "Clear succeeded"}
}
