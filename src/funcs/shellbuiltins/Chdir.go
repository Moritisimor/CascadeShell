package shellbuiltins

import (
	"CaSh/funcs/color"
	"CaSh/funcs/smallhelpers"
	"fmt"
	"os"
)

func Chdir(funcArgs []string) {
	if len(funcArgs) < 2 {
		color.PrintRedln(fmt.Sprintf("Illegal argument count! %s requires 1 argument.", funcArgs[0]))
		color.PrintBlueln(fmt.Sprintf("Example: %s ExampleTargetDirectory", funcArgs[0]))
		return
	}
			
	cdErr := os.Chdir(funcArgs[1]); if cdErr != nil {
		color.PrintRedln(cdErr.Error())
	} else {
		color.PrintGreen("Successfully changed directory to ")
		color.PrintYellowln(smallhelpers.GetCurrentDir())
	}
}