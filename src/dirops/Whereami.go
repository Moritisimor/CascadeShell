package dirops

import (
	"fmt"
	"os"
)

// Talkative version, for interacting with the user.
func Whereami() {
	fmt.Printf("You are in: %s\n", Pwd())
}

// Non-talkative version, gives current working directory and that's it.
func Pwd() string {
	currentDir, dirErr := os.Getwd()
	if dirErr != nil {
		return dirErr.Error()
	} 
	return currentDir
}
