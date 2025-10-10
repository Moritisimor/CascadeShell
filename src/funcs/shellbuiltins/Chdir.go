package shellbuiltins

import (
	"fmt"
	"os"
	"CaSh/funcs/smallhelpers"
)

func Chdir(funcArgs []string) {
	if len(funcArgs) < 2 {
		fmt.Printf("Illegal argument count! %s requires 1 argument.\nExample: %s ExampleTargetDirectory\n", funcArgs[1], funcArgs[1])
		return
	}
			
	cdErr := os.Chdir(funcArgs[1]); if cdErr != nil {
		fmt.Println(cdErr.Error())
	} else {
		fmt.Printf("Successfully changed directory to %s\n", smallhelpers.GetCurrentDir())
	}
}