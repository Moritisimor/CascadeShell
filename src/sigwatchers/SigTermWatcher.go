package sigwatchers

import (
	"fmt"
	"os"
	"os/signal"
	"CaSh/funcs/smallhelpers"
)

func StartSigTermWatcher(activityFlag *bool, username string, host string) chan os.Signal {
	sigTermWatcher := make(chan os.Signal, 1)
	signal.Notify(sigTermWatcher, os.Interrupt) 
	go func() {
		for range(sigTermWatcher) {
			if *activityFlag {
				fmt.Println("\nGot sigterm, quitting current process...")
				break
			} else {
				fmt.Println()
				smallhelpers.Drawprompt(username, host, smallhelpers.GetCurrentDir())
			}
		} 
	}()

	return sigTermWatcher
}
