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
	smallhelpers.FlipBool(processFlag) // Mark process as running

	if runErr != nil {
		color.PrintRedln(runErr.Error())
	} else {
		fmt.Printf("Command %s executed successfully.\n", funcArgs[0])
	}

	smallhelpers.FlipBool(processFlag)
}