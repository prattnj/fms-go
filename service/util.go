package service

import (
	"database/sql"
	"github.com/prattnj/fms-go/dal"
	"github.com/prattnj/fms-go/model"
	"math/rand"
	"time"
)

// Generic 500 response, used in all services
var serverErrorStr = "Internal server error"
var tokenErrorStr = "Error: bad token"
var serverError = model.GenericResponse{Success: false, Message: serverErrorStr}

var defaultIDLength = 32

func generateID(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

func commitAndClose(tx *sql.Tx, db *sql.DB, commit bool) error {
	if commit {
		err := tx.Commit()
		if err != nil {
			return err
		}
	} else {
		err := tx.Rollback()
		if err != nil {
			return err
		}
	}
	err := dal.DbClose(db)
	if err != nil {
		return err
	}
	return nil
}
