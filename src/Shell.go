package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func getCurrentDir() string {
	currentDir, _ := os.Getwd()
	return currentDir
}

func main() {
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

	fmt.Print("<|------------------------------------ (INFO) ------------------------------------|>\n" +
			  "Shell made by Moritisimor. GitHub Repo: https://github.com/Moritisimor/CascadeShell.\n" +
			  "<|------------------------------------ (INFO) ------------------------------------|>\n\n")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("( %s@%s#%s ) -> ", currentUser.Username, currentHost, getCurrentDir())
		rawLine, readingErr := reader.ReadString('\n')
		if readingErr != nil {
			fmt.Printf("FATAL! Reading Line failed! Exiting...\nError: %s", readingErr.Error())
			os.Exit(2)
		}

		formattedLine := strings.Split(strings.TrimSpace(rawLine), " ")
		switch strings.ToLower(formattedLine[0]) {
		default:
			cmd := exec.Command(formattedLine[0], formattedLine[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			runErr:= cmd.Run()

			if runErr != nil {
				fmt.Println(runErr.Error())
			} else {
				fmt.Printf("Command %s executed successfully.\n", formattedLine[0])
			}

		case "cd", "changedir", "changedirectory":
			if len(formattedLine) < 2 {
				fmt.Printf("Illegal argument count! %s requires 1 argument.\nExample: %s ExampleTargetDirectory\n", formattedLine[1], formattedLine[1])
				return
			}
			
			os.Chdir(formattedLine[1])
			fmt.Printf("Successfully changed directory to %s\n", getCurrentDir())

		case "gohome":
			os.Chdir(userHome)
			fmt.Printf("Successfully changed directory to %s\n", getCurrentDir())
			fmt.Println("Welcome home!")

		case "clear", "clearscreen", "clr":
			fmt.Print("\033[H\033[2J\033[3J") // ANSI Escape Code for clearing screen and scrollback buffer

		case "exit", "quit", "bye":
			fmt.Println("Bye!")
			os.Exit(0)

		case "print", "say", "echo":
			printBuf := ""
			for _, i := range formattedLine[1:] {
				printBuf += i + " "
			}
			fmt.Printf("%s\n", printBuf)

		case "":
			// Do nothing.
		}
	}
}
