package fileops

import (
	"fmt"
	"os"
)

func Makefile(fileName string) {
	_, creationErr := os.Create(fileName)
	if creationErr != nil {
		fmt.Println(creationErr.Error())
		return
	}
	fmt.Printf("Successfully created file '%s'\n", fileName)
}