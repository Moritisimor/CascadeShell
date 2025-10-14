package shellbuiltins

import (
	"CaSh/funcs/color"
	"CaSh/funcs/envvargatherers/environment"
	"CaSh/funcs/smallhelpers"
	"os"
)

func Chdir(funcArgs []string) {
	target := environment.GetUser().HomeDir
	if len(funcArgs) > 1 {
		target = funcArgs[1]
	}
			
	cdErr := os.Chdir(target); if cdErr != nil {
		color.PrintRedln(cdErr.Error())
	} else {
		color.PrintGreen("Successfully changed directory to ")
		color.PrintYellowln(smallhelpers.GetCurrentDir())
	}

	if target == environment.GetUser().HomeDir {
		color.PrintYellowln("Welcome home!")
	}
}