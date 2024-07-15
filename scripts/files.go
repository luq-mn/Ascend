package ardent

import (
	"fmt"
    "os"
    "path/filepath"
)

func GetAppDataDirs() (string) {
    appData := os.Getenv("APPDATA")
    
    if appData == "" {
        appData = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Roaming")
    }
    
    return appData
}

func GetFiles(directory string) ([]string) {
	var folders []string
	entries, err := os.ReadDir(directory)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}

	return folders
}

func DataFolder() (string) {
	appDataFolder := GetAppDataDirs()
	saveFolder := filepath.Join(appDataFolder, "Ascend")

	_, err := os.Stat(saveFolder)

	if os.IsNotExist(err) {
		fmt.Println("Ascend data folder does not exist. Creating...")
		err := os.MkdirAll(saveFolder, 0755)
		if err != nil {
			panic(err)
		}
		fmt.Println("Folder created successfully.")

		db := DatabaseConnect(filepath.Join(saveFolder, "test.db"))
		DatabaseExec(db, "CREATE TABLE IF NOT EXISTS groups (name TEXT PRIMARY KEY, description TEXT)")
		DatabaseClose(db)

		return saveFolder

	} else if err != nil {
		panic(err)

	} else {
		return saveFolder

	}
}

func SaveFolder(name string) (string) {
	dataFolder := DataFolder()
	SaveFolder := filepath.Join(dataFolder, name)

	_, err := os.Stat(SaveFolder)

	if os.IsNotExist(err) {
		fmt.Println("Creating folder...")
		err := os.MkdirAll(SaveFolder, 0755)
		if err != nil {
			panic(err)
		}
		fmt.Println("Folder created successfully.")

		return SaveFolder

	} else if err != nil {
		panic(err)

	} else {
		return SaveFolder

	}
}