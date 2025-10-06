package dirops

import (
	"fmt"
	"os"
	"strings"
)

func Lookaround(dirName string) {
	var dir string
	if strings.TrimSpace(dirName) == "" {
		dir = "./"
	}

	dirs, err := os.ReadDir(dir)

	if err != nil {
		fmt.Println()
	} else {
		entryCounter := 0
		fmt.Println("Subdirectories and Files:")
		for _, entry := range(dirs) {
			fmt.Printf("-> %s\n", entry.Name())
			entryCounter++
		}
		if entryCounter == 0 {
			fmt.Println("Empty directory.")
		}
	}
}
