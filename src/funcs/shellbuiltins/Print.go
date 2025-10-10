package shellbuiltins

import (
	"fmt"
	"strings"
)

func Say(funcArgs []string, symbolTable map[string]string) {
	if len(funcArgs) < 2 {
		usedCommand := funcArgs[0]
		fmt.Printf("Illegal argument count! %s requires at least 1 argument." + 
				   "\nExample usage: %s hello" + 
				   "\nTo access declared variables, put an '@' before it." + 
				   "\nExample usage: %s @shell\n",
				   usedCommand, usedCommand, usedCommand)

		return
	}

	printBuf := ""
	for _, i := range funcArgs[1:] {
		if strings.HasPrefix(i, "@") {
			if symbolTable[i] != "" {
				printBuf += symbolTable[i]
			} else {
				printBuf = "Unknown variable: " + i
			}
		} else {
			printBuf += i + " "
		}
	}
	fmt.Printf("%s\n", printBuf)
}
