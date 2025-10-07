package dirops

import (
	"fmt"
	"os"
)

// Returns true or false if operation succeeds
func Chdir(target string) bool {
	err := os.Chdir(target)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Printf("Successfully changed directory to %s\n", Pwd())
	return true
}