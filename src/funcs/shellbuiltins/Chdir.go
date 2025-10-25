package shellbuiltins

import (
	"CaSh/funcs/color"
	"CaSh/funcs/smallhelpers"
	"os"
	"os/user"
)

func Chdir(funcArgs []string) {
	currentUser, _ := user.Current()
	target := currentUser.HomeDir
	if len(funcArgs) > 1 {
		target = funcArgs[1]
	}
			
	cdErr := os.Chdir(target); if cdErr != nil {
		color.PrintRedln(cdErr.Error())
	} else {
		color.PrintGreen("Successfully changed directory to ")
		color.PrintYellowln(smallhelpers.GetCurrentDir())
	}

	if target == currentUser.HomeDir {
		color.PrintYellowln("Welcome home!")
	}
}