package dal

import "database/sql"

func db() *sql.DB {
	database, err := sql.Open("sqlite3", "fms.db")
	if err != nil {
		panic(err)
	}
	return database

}
