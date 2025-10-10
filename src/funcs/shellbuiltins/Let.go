package shellbuiltins

import (
	"fmt"
	"strings"
)

func Let(funcArgs []string, symbolTable map[string]string) {
	if len(funcArgs) < 3 {
				fmt.Println("Illegal argument count! Let requires at least 2 arguments!" + 
							"\nExample Usage: let greeting hello")

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
		fmt.Printf("Variable %s declared successfully.\n", trueVar)
}