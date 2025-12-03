package shellbuiltins

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/CascadeShell/internals/helpers"
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
		color.PrintYellowln(helpers.GetCurrentDir())
	}

	if target == currentUser.HomeDir {
		color.PrintYellowln("Welcome home!")
	}
}