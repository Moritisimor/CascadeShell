package shellbuiltins

import (
	"fmt"
	"strings"
	"os"
)

func Lookaround(funcArgs []string) {
	targetDir := "./"
	if len(funcArgs) > 1 {
		targetDir = strings.TrimSpace(funcArgs[1])
	}

	entries, readingErr := os.ReadDir(targetDir)
	if readingErr != nil {
		fmt.Println(readingErr.Error())
		return
	}

	entryCount := 0
	for _, entry := range(entries) {
		entryType := ""
		if entry.Type().IsDir() {
			entryType += "Directory"
		} else if entry.Type().IsRegular() {
			entryType += "File"
		} else {
			entryType += "Misc"
		}

		fmt.Printf("-> %s (%s)\n", entry.Name(), entryType)
		entryCount++
	}

	if entryCount == 0 {
		fmt.Println("Empty directory.")
	}
}