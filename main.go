package main

import (
	"fmt"
	"ascend/scripts"
	"path/filepath"
    _ "github.com/mattn/go-sqlite3"
)


func main() {
	fmt.Println("ASCEND")

	// Get the path to the save folder
	fmt.Printf("Enter the path to the save folder: ")
	var save_path string = ardent.ConsoleRead()

	ardent.DatabaseConnect(filepath.Join(save_path, "test.db"))
}
