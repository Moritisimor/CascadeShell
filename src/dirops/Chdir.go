package dirops

import (
	"fmt"
	"os"
)

// Returns true or false if operation succeeds
func Chdir(target string) bool {
	err := os.Chdir(target)
	if err != nil {
		fmt.Printf("Could not change directory to %s! Ensure %s exists and is not a file.\n", target, target)
		return false
	}
	fmt.Printf("Successfully changed directory to %s\n", Pwd())
	return true
}