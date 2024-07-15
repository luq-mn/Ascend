package main

import (
	// "os"
	// "strings"
	"fmt"
	"ascend/scripts"
	"path/filepath"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
	ardent.ConsoleClear()
	fmt.Printf("ASCEND\nTasks with scores\n\n")

	dataFolder := ardent.DataFolder()

	// Print out all the folders
	allSaveFolders := ardent.GetFiles(dataFolder)

	if len(allSaveFolders) == 0 {
		fmt.Println("No save folders found.")
	} else {
		fmt.Println("Save folders:")
		for _, folder := range allSaveFolders {
			fmt.Printf("- %s\n", folder)
		}
	}
	fmt.Println("\nEnter name of a listed folder, if it does not exist, we'll create a new one")

	// Enter folder name
	fmt.Printf("Folder name: ")
	saveFolder := ardent.ConsoleRead()
	saveFolder = ardent.SaveFolder(saveFolder)
	
	db:= ardent.DatabaseConnect(filepath.Join(saveFolder, "test.db"))

	// Print out all tasks
	ardent.ConsoleClear()
	fmt.Printf("Folder: %s \n", saveFolder)

	

	// Database things
	ardent.DatabaseExec(db, "CREATE TABLE IF NOT EXISTS test (id INTEGER, users TEXT)")
	ardent.DatabaseExec(db, "INSERT INTO test (id, users) VALUES (1, 'luqman')")
}