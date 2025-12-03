package helpers

import (
	"github.com/chzyer/readline"
	"io"
	"os"
	"log"
)

func GetLine(reader readline.Instance) string {
	rawLine, readingErr := reader.Readline()
	if readingErr != nil {
		switch readingErr {
		case readline.ErrInterrupt:
			rawLine = ""
		case io.EOF:
			os.Exit(0)
		default:
			log.Fatal(readingErr.Error())
		}
	}

	return rawLine
}