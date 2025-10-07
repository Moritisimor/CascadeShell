package dirops

import (
	"fmt"
	"os"
)

func Makedir(dirName string) {
	if os.Mkdir(dirName, 0755) != nil {
		fmt.Printf("Creating directory '%s' failed! " + 
				   "Ensure you have the necessary permissions and that it doesn't exist already.\n", dirName)
		return
	}
	fmt.Printf("Creating directory '%s' succeeded.\n", dirName)
}