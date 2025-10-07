package dirops

import (
	"fmt"
	"os"
)

func Makedir(dirName string) {
	err := os.Mkdir(dirName, 0755); if err != nil {
		fmt.Println()
		return
	}
	fmt.Printf("Creating directory '%s' succeeded.\n", dirName)
}