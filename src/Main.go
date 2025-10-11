package main

import (
	"CaSh/funcs/color"
	"CaSh/funcs/envvargatherers/environment"
	"CaSh/funcs/shellbuiltins"
	"CaSh/funcs/smallhelpers"
	"CaSh/sigwatchers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	currentUser := environment.GetUser()
	currentHost := environment.GetHost()
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
	sigwatchers.StartSigTermWatcher(&activeProcess, currentUser.Username, currentHost)

	color.PrintBlueln("Cascade Shell, Made by Moritisimor.\nhttps://github.com/Moritisimor/CascadeShell\n")

	shellbuiltins.EpsilonFetch()

	reader := bufio.NewReader(os.Stdin)
	for {
		smallhelpers.Drawprompt(currentUser.Username, currentHost, smallhelpers.GetCurrentDir())
		rawLine, readingErr := reader.ReadString('\n')
		if readingErr != nil {
			color.PrintRed(fmt.Sprintf("FATAL! Reading Line failed! Exiting...\nError: %s", readingErr.Error()))
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
			color.PrintGreenln(fmt.Sprintf("Successfully changed directory to %s", smallhelpers.GetCurrentDir()))
			color.PrintYellowln("Welcome home!")

		case "let", "var":
			shellbuiltins.Let(formattedLine, defVars)

		case "unlet", "free":
			shellbuiltins.Unlet(formattedLine, defVars)

		case "lookaround", "ls":
			shellbuiltins.Lookaround(formattedLine)

		case "whereami", "pwd":
			color.PrintBlue("You are in: ")
			color.PrintYellowln(smallhelpers.GetCurrentDir())

		case "whoami":
			color.PrintBlue("You are: ")
			color.PrintGreenln(currentUser.Username)

		case "clear", "clearscreen", "clr":
			fmt.Print("\033[H\033[2J\033[3J") // ANSI Escape Code for clearing screen and scrollback buffer

		case "exit", "quit", "bye":
			color.PrintBlueln("Bye!")
			os.Exit(0)

		case "help":
			color.PrintBlueln("I was too lazy to implement a real help command so visit https://github.com/Moritisimor/CascadeShell")

		case "print", "say":
			shellbuiltins.Say(formattedLine, defVars)

		case "epsilon", "epsilonfetch":
			shellbuiltins.EpsilonFetch()

		case "":
			// Do nothing.
		}

		defVars["@lastcommand"] = strings.Join(formattedLine[0:], " ")
	}
}
