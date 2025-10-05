package fileops

import (
	"os"
)

func Writefile(fileName string, fileContent string) {
	err := os.WriteFile(fileName, []byte(fileContent), os.ModeAppend)
	if err != nil {
		
	} else {

	}
}
