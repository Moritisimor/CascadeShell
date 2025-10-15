package smallhelpers

import (
	"CaSh/funcs/envvargatherers/environment"
	"os"
	"strings"
)

func GetCurrentDir() string {
	currentDir, _ := os.Getwd()
	if strings.HasPrefix(currentDir, environment.GetUser().HomeDir) {
		return "~" + strings.TrimPrefix(currentDir, environment.GetUser().HomeDir)
	}
	return currentDir
}