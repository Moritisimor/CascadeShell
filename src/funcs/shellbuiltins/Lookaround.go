package shellbuiltins

import (
	"CaSh/funcs/color"
	"os"
	"strings"
)

func Lookaround(funcArgs []string) {
	showhidden := false
	for _, i := range(funcArgs) {
		if strings.Contains(i, "--showhidden") {
			showhidden = true
		}
	}

	targetDir := "./"
	if len(funcArgs) > 1 {
		if !strings.HasPrefix(strings.TrimSpace(funcArgs[1]), "--") {
			targetDir = strings.TrimSpace(funcArgs[1])
		}
	}

	entries, readingErr := os.ReadDir(targetDir)
	if readingErr != nil {
		color.PrintRedln(readingErr.Error())
		return
	}

	entryCount := 0
	for _, entry := range(entries) {
		if !showhidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		var entryType string
		if entry.Type().IsDir() {
			entryType = "Directory"
		} else if entry.Type().IsRegular() {
			entryType = "File"
		} else {
			entryType = "Misc"
		}

		color.PrintBlue("-> ")
		color.PrintYellow(entry.Name())
		color.PrintBlue(" (")
		switch entryType { // Perhaps not the most efficient, but it'll do for now.
		default:
			color.PrintMagenta(entryType)
		case "File":
			color.PrintGreen(entryType)
		case "Directory":
			color.PrintRed(entryType)
		}
		color.PrintBlueln(")")
		entryCount++
	}

	if entryCount == 0 {
		color.PrintBlueln("Empty directory.")
	}
}