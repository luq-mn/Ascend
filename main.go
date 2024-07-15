package main

import (
	"strings"
	"fmt"
	"ascend/scripts"
	"path/filepath"
    _ "github.com/mattn/go-sqlite3"
)


func main() {
	fmt.Println("ASCEND")

	// Get the path to the save folder
	fmt.Printf("Enter the path to the save folder: ")
	save_path := strings.TrimSpace(ardent.ConsoleRead())
	
	db:= ardent.DatabaseConnect(filepath.Join(save_path, "test.db"))

	ardent.DatabaseExec(db, "CREATE TABLE IF NOT EXISTS test (id INTEGER, users TEXT)")
	ardent.DatabaseExec(db, "INSERT INTO test (id, users) VALUES (1, 'luqman')")
}