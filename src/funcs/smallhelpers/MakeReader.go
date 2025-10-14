package smallhelpers

import (
	"github.com/chzyer/readline"
	"CaSh/funcs/envvargatherers/environment"
	"log"
)

func MakeReader() readline.Instance {
	reader, readerCreateErr := readline.NewEx(&readline.Config {
			Prompt: SPrintPrompt(environment.GetUser().Username, environment.GetHost(), GetCurrentDir()),
			InterruptPrompt: "^C",
			HistoryFile: GetHistory(),
		})

	if readerCreateErr != nil {
		log.Fatal(readerCreateErr.Error())
	}

	return *reader
}
