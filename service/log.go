package service

import (
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
	"log"
	"time"
)

func Log(endpoint string, ipv4 string, success bool) model.GenericResponse {

	db := dal.Db()
	if db == nil {
		logError("Could not connect to database")
		return serverError
	}
	tx, err := db.Begin()
	if err != nil {
		logError(err.Error())
		return serverError
	}

	logObj := model.Log{Timestamp: time.Now().Unix(), Endpoint: endpoint, IPv4: ipv4, Success: success}
	err = dal.L_insert(tx, logObj)
	if err != nil {
		logError(err.Error())
		err := tx.Rollback()
		if err != nil {
			return serverError
		}
		return serverError
	}

	err = tx.Commit()
	if err != nil {
		logError(err.Error())
		return serverError
	}
	err = dal.DbClose(db)
	if err != nil {
		logError(err.Error())
		return serverError
	}
	return model.GenericResponse{Success: true, Message: "Log succeeded"}

}

func logError(err string) {
	if err != "" {
		log.Printf("FMS-GO LOGGING ERROR: %s\n", err)
	}
}
