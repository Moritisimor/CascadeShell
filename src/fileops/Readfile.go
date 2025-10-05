package fileops

import (
	"fmt"
	"os"
)

func Readfile(fileName string) {
	openedFile, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Opening file failed! Make sure that the file you specified exists!")
		return
	}

	fmt.Printf("File content: %s", string(openedFile))
}