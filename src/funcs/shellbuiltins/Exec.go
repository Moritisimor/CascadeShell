package shellbuiltins

import (
	"CaSh/funcs/color"
	"CaSh/funcs/smallhelpers"
	"fmt"
	"os"
	"os/exec"
)

func Execute(funcArgs []string, processFlag *bool) {
	cmd := exec.Command(funcArgs[0], funcArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	runErr:= cmd.Run()
	smallhelpers.SetTrue(processFlag) // Mark process as running
	defer smallhelpers.SetFalse(processFlag) // Mark as done at the end of the scope

	if runErr != nil {
		color.PrintRedln(runErr.Error())
	} else {
		color.PrintGreenln(fmt.Sprintf("Command %s executed successfully.\n", funcArgs[0]))
	}
}