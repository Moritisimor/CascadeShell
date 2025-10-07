package fileops

import (
	"fmt"
	"os"
)

func Readfile(fileName string) {
	openedFile, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("File content: \n%s\n", string(openedFile))
}