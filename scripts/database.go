package ardent

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// DatabaseConnect opens a connection to a sqlite3 database.
func DatabaseConnect(location string) *sql.DB {
	db, err := sql.Open("sqlite3", location)
	if err != nil {
		panic(err)
	}

	return db
}

func DatabaseExec(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}