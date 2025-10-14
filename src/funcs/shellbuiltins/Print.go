package shellbuiltins

import (
	"CaSh/funcs/color"
	"fmt"
	"strings"
)

func Say(funcArgs []string, symbolTable map[string]string) {
	if len(funcArgs) < 2 {
		color.PrintRedln(fmt.Sprintf("Illegal argument count! %s requires at least 1 argument.", funcArgs[0]))
		color.PrintBlueln(fmt.Sprintf("Example usage: %s hello", funcArgs[0]))
		color.PrintGreenln("To access declared variables, put an '@' before it.")
		color.PrintBlueln(fmt.Sprintf("Example usage: %s @shell", funcArgs[0]))

		return
	}

	printBuf := ""
	for _, i := range funcArgs[1:] {
		if strings.HasPrefix(i, "@") {
			if symbolTable[i] != "" {
				printBuf += symbolTable[i] + " "
			} else {
				printBuf = "Unknown variable: " + i
			}
		} else {
			printBuf += i + " "
		}
	}
	fmt.Printf("%s\n", printBuf)
}
