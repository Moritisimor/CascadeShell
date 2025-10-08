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

	defVars := map[string]string {
		"@shell": 		"cash",
		"@user":  		currentUser.Name,
		"@userID":		currentUser.Uid,
		"@host":		currentHost,
		"@home":		userHome,
		"@lastcommand": "undefined",
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

		case "cd", "chdir":
			if len(formattedLine) < 2 {
				fmt.Printf("Illegal argument count! %s requires 1 argument.\nExample: %s ExampleTargetDirectory\n", formattedLine[1], formattedLine[1])
				return
			}
			
			cdErr := os.Chdir(formattedLine[1]); if cdErr != nil {
				fmt.Println(cdErr.Error())
			} else {
				fmt.Printf("Successfully changed directory to %s\n", getCurrentDir())
			}

		case "gohome":
			os.Chdir(userHome)
			fmt.Printf("Successfully changed directory to %s\n", getCurrentDir())
			fmt.Println("Welcome home!")

		case "let":
			if len(formattedLine) < 3 {
				fmt.Println("Illegal argument count! Let requires at least 2 arguments!" + 
							"\nExample Usage: let greeting hello")
			} else {
				trueVar := ""
				if strings.HasPrefix(formattedLine[1], "@") {
					trueVar = formattedLine[1]
				} else {
					trueVar = "@" + formattedLine[1]
				}

				starter := 2
				if formattedLine[2] == "=" {
					starter = 3
				}

				defVars[trueVar] = strings.Join(formattedLine[starter:], " ")
				fmt.Printf("Variable %s declared successfully.\n", trueVar)
			}

		case "unlet": // This implementation of undeclaring a variable is more convention-based than anything else, but good enough.
			if len(formattedLine) < 2 {
				fmt.Println("Illegal argument count! Unlet requires at least 2 arguments!" + 
							"\nExample usage: unlet hello")
			} else {
				trueVar := formattedLine[1]
				if !strings.HasPrefix(trueVar, "@") {
					trueVar = "@" + trueVar
				}

				if defVars[trueVar] == "" || defVars[trueVar] == "nil" {
					fmt.Printf("%s does not exist.\n", trueVar)
				} else {
					delete(defVars, trueVar)
					fmt.Printf("%s uninitialized successfully.\n", trueVar)
				}
			}

		case "lookaround", "ls": // A rather stupid ls implementation since it also lists entries which are hidden (starting with .) but good enough as a basic util.
			targetDir := "./"
			if len(formattedLine) < 1 {
				targetDir += strings.TrimSpace(formattedLine[1])
			}

			entries, readingErr := os.ReadDir(targetDir)
			if readingErr != nil {
				fmt.Println(readingErr.Error())
				return
			}

			entryCount := 0
			for _, entry := range(entries) {
				entryType := ""
				if entry.Type().IsDir() {
					entryType += "Directory"
				} else if entry.Type().IsRegular() {
					entryType += "File"
				} else {
					entryType += "Misc"
				}

				fmt.Printf("-> %s (%s)\n", entry.Name(), entryType)
				entryCount++
			}

			if entryCount == 0 {
				fmt.Println("Empty directory.")
			}

		case "whereami", "pwd":
			fmt.Printf("You are in: %s\n", getCurrentDir())

		case "clear", "clearscreen", "clr":
			fmt.Print("\033[H\033[2J\033[3J") // ANSI Escape Code for clearing screen and scrollback buffer

		case "exit", "quit", "bye":
			fmt.Println("Bye!")
			os.Exit(0)

		case "print", "say":
			if len(formattedLine) < 2 {
				fmt.Printf("Illegal argument count! %s requires at least 1 argument.", formattedLine[0] + 
							"\nExample usage: print hello" + 
							"\nTo access declared variables, put an '@' before it." + 
							"\nExample usage: print @shell")
			} else {
				printBuf := ""
				for _, i := range formattedLine[1:] {
					if defVars[i] != "" {
						printBuf += defVars[i] + " "
					} else {
						printBuf += i + " "
					}
				}
				fmt.Printf("%s\n", printBuf)
			}

		case "":
			// Do nothing.
		}
		defVars["@lastcommand"] = strings.Join(formattedLine[0:], " ")
	}
}
