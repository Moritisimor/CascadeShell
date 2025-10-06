package main

import (
	"CaSh/dirops"
	// "CaSh/fileops"
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	// !----------- Start Gathering environment variables -----------!
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
	// !----------- Stop Gathering environment variables -----------!

	// !----------- Start REPL -----------!
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("[ %s@%s#%s ] -> ", currentUser.Username, currentHost, currentDir)
		rawLine, readingErr := reader.ReadString('\n')
		if readingErr != nil {
			fmt.Printf("(FATAL) Reading Line failed! Exiting...\nError: %e", readingErr)
			os.Exit(2)
		}

		formattedLine := strings.Split(strings.TrimSpace(rawLine), " ")
		switch strings.ToLower(formattedLine[0]) {
		default:
			fmt.Printf("Unknown command: %s\n", formattedLine[0])

		case "lookaround":
			if len(formattedLine) == 1 {
				dirops.Lookaround("")
			} else {
				dirops.Lookaround(formattedLine[1])
			}

		case "whereami":
			dirops.Whereami()

		case "chdir", "changedir", "changedirectory":
			if len(formattedLine) == 1 {
				dirops.Chdir(currentHome)
			} else {
				dirops.Chdir(formattedLine[1])
			}
			currentDir = dirops.Pwd()

		case "gohome":
			dirops.Chdir(currentHome)

		case "exit", "quit":
			fmt.Println("Bye!")
			os.Exit(0)

		case "print", "say", "echo":
			printBuf := ""
			for _, i := range(formattedLine[1:]) {
				printBuf += i + " "
			}
			fmt.Printf("%s\n", printBuf)
		}
	}
	// !----------- Stop REPL -----------!
}
