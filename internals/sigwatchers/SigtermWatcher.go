package sigwatchers

import (
	"fmt"
	"os"
	"os/signal"
)

func StartSigTermWatcher(activityFlag *bool, username string, host string)  {
	sigTermWatcher := make(chan os.Signal, 1)
	signal.Notify(sigTermWatcher, os.Interrupt) 
	go func() {
		for range(sigTermWatcher) {
			if *activityFlag {
				fmt.Println("\nGot sigterm, quitting current process...")
				break
			} else {
				fmt.Println()
			}
		} 
	}()
}