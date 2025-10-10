package shellbuiltins

import (
	"fmt"
	"os"
	"os/exec"
	"CaSh/funcs/smallhelpers"
)

func Execute(funcArgs []string, processFlag *bool) {
	cmd := exec.Command(funcArgs[0], funcArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	runErr:= cmd.Run()
	smallhelpers.FlipBool(processFlag)

	if runErr != nil {
		fmt.Println(runErr.Error())
	} else {
		fmt.Printf("Command %s executed successfully.\n", funcArgs[0])
	}

	smallhelpers.FlipBool(processFlag)
}