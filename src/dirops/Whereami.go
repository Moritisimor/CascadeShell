package dirops

import (
	"os"
	"fmt"
)

func Whereami() {
	currentDir, dirErr := os.Getwd()
	if dirErr != nil {
		fmt.Println("I don't know.")
		return
	} 
	fmt.Printf("You are in: %s\n", currentDir)
}
