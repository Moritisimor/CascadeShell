package helpers

import (
	"log"
	"os"
	"os/user"

	"github.com/chzyer/readline"
)

func MakeReader() readline.Instance {
	currentUser, _ := user.Current()
	currentHost, _ := os.Hostname()
	reader, readerCreateErr := readline.NewEx(&readline.Config {
		Prompt: SPrintPrompt(currentUser.Username, currentHost, GetCurrentDir()),
		InterruptPrompt: "^C",
		HistoryFile: GetHistory(),
	})

	if readerCreateErr != nil {
		log.Fatal(readerCreateErr.Error())
	}

	return *reader
}
