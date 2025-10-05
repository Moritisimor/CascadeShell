package fileops

import (
	"fmt"
	"os"
)

func Makefile(fileName string) {
	_, creationErr := os.Create(fileName)
	if creationErr != nil {
		fmt.Printf("Could not create file!\nError: %e\n", creationErr)
		return
	}
	fmt.Printf("Successfully created file '%s'", fileName)
}