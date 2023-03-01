package dal

import (
	"database/sql"
	"github.com/prattnj/fms-go/model"
)

func L_insert(tx *sql.Tx, log model.Log) error {
	_, err := tx.Exec("INSERT INTO log (timestamp, endpoint, ipv4, success) VALUES(?,?,?,?);", log.Timestamp, log.Endpoint, log.IPv4, log.Success)
	if err != nil {
		return err
	}
	return nil
}
