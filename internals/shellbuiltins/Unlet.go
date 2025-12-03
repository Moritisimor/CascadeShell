package shellbuiltins

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"fmt"
	"strings"
)

func Unlet(funcArgs []string, symbolTable map[string]string) {
	if len(funcArgs) < 2 {
		color.PrintRedln(fmt.Sprintf("Illegal argument count! %s requires 1 argument!", funcArgs[0]))
		color.PrintBlueln(fmt.Sprintf("Example usage: %s @hello", funcArgs[0]))

		return
	} 
	
	trueVar := funcArgs[1]
	if !strings.HasPrefix(trueVar, "@") {
		trueVar = "@" + trueVar
	}

	if symbolTable[trueVar] == "" || symbolTable[trueVar] == "nil" {
		color.PrintMagenta(trueVar)
		color.PrintRedln(" does not exist.")
		return
	}

	delete(symbolTable, trueVar)
	color.PrintMagenta(trueVar)
	color.PrintGreenln(" has been uninitialized successfully.")
}
