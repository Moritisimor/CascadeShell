package shellbuiltins

import (
	"strings"
	"fmt"
)

func Unlet(funcArgs []string, symbolTable map[string]string) {
	if len(funcArgs) < 2 {
		fmt.Println("Illegal argument count! Unlet requires at least 2 arguments!" + 
					"\nExample usage: unlet hello")

		return
	} 
	
	trueVar := funcArgs[1]
	if !strings.HasPrefix(trueVar, "@") {
		trueVar = "@" + trueVar
	}
	if symbolTable[trueVar] == "" || symbolTable[trueVar] == "nil" {
		fmt.Printf("%s does not exist.\n", trueVar)
	} else {
		delete(symbolTable, trueVar)
		fmt.Printf("%s uninitialized successfully.\n", trueVar)
	}
}
