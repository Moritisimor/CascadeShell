package shellbuiltins

import (
	"CaSh/funcs/color"
	"fmt"
	"strings"
)

func Let(funcArgs []string, symbolTable map[string]string) {
	if len(funcArgs) < 3 {
		color.PrintRedln(fmt.Sprintf("Illegal argument count! %s requires at least 2 arguments!", funcArgs[0]))
		color.PrintBlueln(fmt.Sprintf("Example Usage: %s greeting = hello", funcArgs[0]))

		return
	}
	trueVar := ""
	if strings.HasPrefix(funcArgs[1], "@") {
		trueVar = funcArgs[1]
	} else {
		trueVar = "@" + funcArgs[1]
	}

	starter := 2
	if funcArgs[2] == "=" {
		starter = 3
	}

	symbolTable[trueVar] = strings.Join(funcArgs[starter:], " ")
	color.PrintGreenln(fmt.Sprintf("Variable %s declared successfully.", trueVar))
}