package dirops

import (
	"os"
	"fmt"
)

func Whereami() {
	currentDir, dirErr := os.Getwd()
	if dirErr == nil {
		fmt.Printf("You are in: %s\n", currentDir)
	} else {
		fmt.Println("I don't know.")
	}
}
