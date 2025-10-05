package fileops

import (
	"fmt"
	"os"
)

func Writefile(fileName string, fileContent string) {
	err := os.WriteFile(fileName, []byte(fileContent), os.ModeAppend)
	if err != nil {
		fmt.Println("Writing file failed! ensure you have the neceesary permissions to write to the specified directory!")
		return
	}
	fmt.Printf("Successfully wrote specified content to %s\n", fileName)
}
