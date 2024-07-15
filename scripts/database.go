package ardent

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// DatabaseConnect opens a connection to a SQLite database.
func DatabaseConnect(location string) *sql.DB {
	db, err := sql.Open("sqlite3", location)
	if err != nil {
		panic(err)
	}

	return db
}

// Executes a query to the database
func DatabaseExec(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func DatabaseQuery(db *sql.DB, query string) *sql.Rows {
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	return rows
}

func DatabaseClose(db *sql.DB) {
	db.Close()
}