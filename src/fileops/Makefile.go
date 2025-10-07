package fileops

import (
	"fmt"
	"os"
)

func Makefile(fileName string) {
	_, creationErr := os.Create(fileName)
	if creationErr != nil {
		fmt.Printf("Could not create file! Make sure %s does not already exist.\n", fileName)
		return
	}
	fmt.Printf("Successfully created file '%s'\n", fileName)
}