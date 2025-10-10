package main

import (
	"CaSh/funcs/envvargatherers"
	"CaSh/funcs/shellbuiltins"
	"CaSh/funcs/smallhelpers"
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

func main() {
	currentUser := envvargatherers.GetUser()
	currentHost := envvargatherers.GetHome()
	userHome := currentUser.HomeDir
	

	os.Chdir(userHome)

	defVars := map[string]string {
		"@shell": 		"cash",
		"@user":  		currentUser.Name,
		"@userID":		currentUser.Uid,
		"@host":		currentHost,
		"@home":		userHome,
		"@lastcommand": "undefined",
	}

	activeProcess := false

	sigTermWatcher := make(chan os.Signal, 1)
		signal.Notify(sigTermWatcher, os.Interrupt) 
		go func() {
			for range(sigTermWatcher) {
				if activeProcess {
					fmt.Println("\nGot sigterm, quitting current process...")
					break
				} else {
					fmt.Printf("\n( %s@%s#%s ) -> ", currentUser.Username, currentHost, smallhelpers.GetCurrentDir())
				}
			} 
		}()

	fmt.Print("<|------------------------------------ (INFO) ------------------------------------|>\n" +
			  "Shell made by Moritisimor. GitHub Repo: https://github.com/Moritisimor/CascadeShell.\n" +
			  "<|------------------------------------ (INFO) ------------------------------------|>\n\n")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("( %s@%s#%s ) -> ", currentUser.Username, currentHost, smallhelpers.GetCurrentDir())
		rawLine, readingErr := reader.ReadString('\n')
		if readingErr != nil {
			fmt.Printf("FATAL! Reading Line failed! Exiting...\nError: %s", readingErr.Error())
			os.Exit(2)
		}

		formattedLine := strings.Split(strings.TrimSpace(rawLine), " ")
		switch strings.ToLower(formattedLine[0]) {
		default:
			shellbuiltins.Execute(formattedLine, &activeProcess)

		case "cd", "chdir":
			shellbuiltins.Chdir(formattedLine)

		case "gohome":
			os.Chdir(userHome)
			fmt.Printf("Successfully changed directory to %s\n", smallhelpers.GetCurrentDir())
			fmt.Println("Welcome home!")

		case "let":
			shellbuiltins.Let(formattedLine, defVars)

		case "unlet":
			shellbuiltins.Unlet(formattedLine, defVars)

		case "lookaround", "ls": // A rather stupid ls implementation since it also lists entries which are hidden (starting with .) but good enough as a basic util.
			shellbuiltins.Lookaround(formattedLine)

		case "whereami", "pwd":
			fmt.Printf("You are in: %s\n", smallhelpers.GetCurrentDir())

		case "clear", "clearscreen", "clr":
			fmt.Print("\033[H\033[2J\033[3J") // ANSI Escape Code for clearing screen and scrollback buffer

		case "exit", "quit", "bye":
			fmt.Println("Bye!")
			os.Exit(0)

		case "print", "say":
			shellbuiltins.Say(formattedLine, defVars)

		case "":
			// Do nothing.
		}

		defVars["@lastcommand"] = strings.Join(formattedLine[0:], " ")
	}
}
