package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
	"time"
)

func Log(endpoint string, ipv4 string, success bool) model.GenericResponse {

	db := dal.Db()
	if db == nil {
		return serverError
	}
	tx, err := db.Begin()
	if err != nil {
		return serverError
	}

	log := model.Log{Timestamp: time.Now().Unix(), Endpoint: endpoint, IPv4: ipv4, Success: success}
	err = dal.L_insert(tx, log)
	if err != nil {
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
	return model.GenericResponse{Success: true, Message: "Log succeeded"}

}
