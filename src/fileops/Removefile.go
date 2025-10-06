package fileops

import (
	"os"
	"fmt"
)

func Removefile(target string) {
	if os.Remove(target) != nil {
		fmt.Println("Deleting specified file failed! Ensure that the file exists and that you have the necessary rights." + 
					"\nIn case you specified a directory instead of a file, make sure it is empty or use removedir instead.")
		return
	}
	fmt.Printf("Successfully deleted %s.\n", target)
}
