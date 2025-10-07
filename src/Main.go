package main

import (
	"CaSh/dirops"
	"CaSh/fileops"
	"bufio"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func main() {
	// !----------- Start Gathering environment variables -----------! \\
	currentUser, userErr := user.Current()
	if userErr != nil {
		fmt.Printf("Could not read current user!\nError: %e\n", userErr)
		os.Exit(1)
	}

	userHome := currentUser.HomeDir
	os.Chdir(userHome)

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
	// !----------- Stop Gathering environment variables -----------! \\

	// !----------- Start REPL -----------! \\
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("( %s@%s#%s ) -> ", currentUser.Username, currentHost, currentDir)
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
			showHidden := false
			if strings.Contains(rawLine, "--showhidden") {
				showHidden = true
			}
			
			if len(formattedLine) == 1 {
				dirops.Lookaround("", showHidden) // An empty string means the current directory.
			} else {
				dirops.Lookaround(formattedLine[1], showHidden)
			}

		case "whereami":
			dirops.Whereami()

		case "chdir", "changedir", "changedirectory":
			if len(formattedLine) == 1 {
				dirops.Chdir(userHome)
			} else {
				dirops.Chdir(formattedLine[1])
			}

			if formattedLine[1] == "/" {
				fmt.Println("Be careful when you're in the root directory!")
			}
			currentDir = dirops.Pwd()

		case "mkdir", "makedirectory", "makedir":
			if len(formattedLine) == 1 {
				fmt.Println("Illegal argument count! mkdir requires one following argument, specifying the directory name.\n" + 
							"Example: mdkir ExampleFolder")
			} else {
				dirops.Makedir(formattedLine[1])
			}

		case "makefile":
			if len(formattedLine) < 2 {
				fmt.Println("Illegal Argument Count! To use makefile you must specify the name of the file you want to create." + 
							"Example: makefile hello.txt")
				return
			}
			fileops.Makefile(formattedLine[1])

		case "readfile":
			if len(formattedLine) < 2 {
				fmt.Println("Illegal Argument Count! To use readfile you must specify the name of the file you want to read." + 
							"Example: readfile hello.txt")
				return
			}
			fileops.Readfile(formattedLine[1])

		case "removefile":
			if len(formattedLine) < 2 {
				fmt.Println("Illegal Argument Count! To use removefile you must specify the name of the file you want to remove." + 
							"Example: removefile hello.txt")
				return
			}
			fileops.Removefile(formattedLine[1])

		case "gohome":
			dirops.Chdir(userHome)

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
	// !----------- Stop REPL -----------! \\
}
