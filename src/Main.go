package main

import (
	"fmt"
	"os"
	"os/user"
	"CaSh/fileops"
	"CaSh/dirops"
)

func main() {
	// Gather environment variables
	currentUser, userErr := user.Current()
	if userErr != nil {
		fmt.Printf("Could not read current user!\nError: %e\n", userErr)
		os.Exit(1)
	}

	currentHome := currentUser.HomeDir
	os.Chdir(currentHome)

	currentHost, hostErr := os.Hostname()
	if hostErr != nil {
		fmt.Printf("Could not read OS Hostname!\nError: %e\n", hostErr)
		os.Exit(1)
	}

	currentDir, dirErr := os.Getwd()
	if dirErr != nil {
		fmt.Printf("Could not read current directory!\nError: %e\n", dirErr)
		os.Exit(1)
	}

	fmt.Printf("[ %s@%s#%s ] -> ", currentUser.Username, currentHost, currentDir)
	dirops.Whereami()
	fileops.Makefile("Hello.txt")
	fileops.Readfile("Test.txt")
}
