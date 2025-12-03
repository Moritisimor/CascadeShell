package main

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/CascadeShell/internals/shellbuiltins"
	"github.com/Moritisimor/CascadeShell/internals/helpers"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/Moritisimor/EpsilonFetch/pkg/epsilonfetch"
)

func main() {
	currentUser, userErr := user.Current()
	if userErr != nil {
		color.PrintRedln(userErr.Error())
		return
	}

	currentHost, hostErr := os.Hostname()
	if hostErr != nil {
		color.PrintRedln(hostErr.Error())
		return
	}

	userHome := currentUser.HomeDir

	defVars := map[string]string {
		"@shell": 		"cash",
		"@user":  		currentUser.Name,
		"@userID":		currentUser.Uid,
		"@host":		currentHost,
		"@home":		userHome,
	}

	activeProcess := false

	color.PrintBlueln("Cascade Shell, Made by Moritisimor.\nhttps://github.com/Moritisimor/CascadeShell\n")
	epsilonfetch.EpsilonFetch()
	helpers.MakeHistory()

	for {
		reader := helpers.MakeReader() 
		rawLine := helpers.GetLine(reader)

		for _, subcommand := range(strings.Split(strings.TrimSpace(rawLine), ";")) {
			formattedLine := strings.Split(strings.TrimSpace(subcommand), " ")
			switch strings.ToLower(formattedLine[0]) {
			default:
				shellbuiltins.Execute(formattedLine, &activeProcess)

			case "exec":
				shellbuiltins.Execute(formattedLine[1:], &activeProcess)

			case "cd", "chdir":
				shellbuiltins.Chdir(formattedLine)

			case "gohome":
				os.Chdir(userHome)
				color.PrintGreenln(fmt.Sprintf("Successfully changed directory to %s", helpers.GetCurrentDir()))
				color.PrintYellowln("Welcome home!")

			case "let", "var":
				shellbuiltins.Let(formattedLine, defVars)

			case "showhist", "hist":
				helpers.ReadHistory()

			case "clearhist":
				helpers.ClearHistory()

			case "unlet", "free":
				shellbuiltins.Unlet(formattedLine, defVars)

			case "lookaround", "ls":
				shellbuiltins.Lookaround(formattedLine)

			case "whereami", "pwd":
				color.PrintBlue("You are in: ")
				color.PrintYellowln(helpers.GetCurrentDir())

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
				epsilonfetch.EpsilonFetch()

			case "": // Do nothing.
			}
		}
	}
}
