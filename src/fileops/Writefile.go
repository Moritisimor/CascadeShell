package fileops

import (
	"fmt"
	"os"
)

func Writefile(fileName string, fileContent string) {
	err := os.WriteFile(fileName, []byte(fileContent), 0755)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Successfully wrote specified content to %s\n", fileName)
}
