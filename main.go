package main

import (
    "database/sql"
	"path/filepath"
    _ "github.com/mattn/go-sqlite3"
)


func main() {
	var save_path string = "./ascend"

	db, err := sql.Open("sqlite3", filepath.Join(save_path, "test.db"))
	if err!= nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("INSERT INTO test (name) VALUES (?)", "sarah")
	if err != nil {
		panic(err)
	}


	defer db.Close()
}
