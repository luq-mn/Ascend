package ardent

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// DatabaseConnect opens a connection to a sqlite3 database.
func DatabaseConnect(location string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", location)
	if err!= nil {
		panic(err)
	}

	return db, nil
}