package main

import (
	"os"
	"strings"
	"fmt"
	"ascend/scripts"
	"path/filepath"
    _ "github.com/mattn/go-sqlite3"
)

var save_path string
func main() {
	ardent.ConsoleClear()
	fmt.Println("ASCEND")

	fmt.Println("[1] Connect to a save folder\n[2] Create a new one")
	fmt.Printf("Select an option [1/2]: ")
	option := strings.TrimSpace(ardent.ConsoleRead())
	switch option {
	case "1":
		// Get the path to the save folder
		fmt.Printf("Path to the save folder: ")
		save_path = strings.TrimSpace(ardent.ConsoleRead())

	case "2":
		// Create the folder
		fmt.Printf("Folder name: ")
		folder_name := strings.TrimSpace(ardent.ConsoleRead())
		fmt.Printf("Location: ")
		folder_location := strings.TrimSpace(ardent.ConsoleRead())

		err := os.MkdirAll(filepath.Join(folder_location, folder_name), 0755)
		if err != nil {
			panic(fmt.Sprintf("Failed to create folder: %v", err))
		}
		fmt.Println("Folder created successfully.")

		save_path = filepath.Join(folder_location, folder_name)

	default:
		panic("Invalid option. Please select either 1 or 2.")
	}
	
	db:= ardent.DatabaseConnect(filepath.Join(save_path, "test.db"))

	ardent.DatabaseExec(db, "CREATE TABLE IF NOT EXISTS test (id INTEGER, users TEXT)")
	ardent.DatabaseExec(db, "INSERT INTO test (id, users) VALUES (1, 'luqman')")
}