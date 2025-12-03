package shellbuiltins

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"github.com/Moritisimor/CascadeShell/internal/helpers"
	"fmt"
	"math"
	"os"
	"os/exec"
	"time"
)

func Execute(funcArgs []string, processFlag *bool) {
	startTime := time.Now()
	cmd := exec.Command(funcArgs[0], funcArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	helpers.SetTrue(processFlag) // Mark process as running
	runErr := cmd.Run()
	defer helpers.SetFalse(processFlag) // Mark as done at the end of the scope

	if runErr != nil {
		color.PrintRedln(runErr.Error())
	} else {
		difference := math.Abs(float64(time.Until(startTime)) / 1000000000)
		color.PrintGreenln(fmt.Sprintf("\nCommand %s executed successfully in %.2f second/s.", funcArgs[0], difference))
	}
}
