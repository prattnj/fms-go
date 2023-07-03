package dal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func Db() *sql.DB {
	password := os.Getenv("MYSQL_PASSWORD")
	if password == "" {
		return nil
	}
	database, err := sql.Open("mysql", "pratt:"+password+"@tcp(localhost:3306)/fms")
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return nil
	}
	err = database.Ping()
	if err != nil {
		fmt.Println("Error pinging database: ", err)
		return nil
	}
	return database
}

func DbClose(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
