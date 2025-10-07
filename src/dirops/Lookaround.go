package dirops

import (
	"fmt"
	"os"
	"strings"
)

func Lookaround(dirName string, showhidden bool) {
	dirs, err := os.ReadDir("./" + dirName)

	if err != nil {
		fmt.Printf("Something went while reading the entries of %s. Make sure %s is a folder and exists.\n", dirName, dirName)
		return
	}

	entryCounter := 0
	fmt.Println("Subdirectories and Files:")
	for _, entry := range(dirs) {
		entryCounter++
		entryType := ""
		if entry.Type().IsDir() {
			entryType = "Folder"
		} else if entry.Type().IsRegular() {
			entryType = "File"
		} else {
			entryType = "Misc"
		}

		if !showhidden {
			if !strings.HasPrefix(entry.Name(), ".") {
				fmt.Printf("-> %s (%s)\n", entry.Name(), entryType)
			}
		} else {
			fmt.Printf("-> %s (%s)\n", entry.Name(), entryType)
		}
	}
		
	if entryCounter == 0 {
		fmt.Println("Empty directory.")
	}
}
