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
	fmt.Println("Subdirectories:")

	if err != nil {
		fmt.Println()
	} else {
		for _, i := range(dirs) {
			fmt.Printf("-> %s\n", i.Name())
		}
	}
}
