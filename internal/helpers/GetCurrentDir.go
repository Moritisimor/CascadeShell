package helpers

import (
	"os"
	"os/user"
	"strings"
)

func GetCurrentDir() string {
	currentUser, _ := user.Current()
	currentDir, _ := os.Getwd()
	if strings.HasPrefix(currentDir, currentUser.HomeDir) {
		return "~" + strings.TrimPrefix(currentDir, currentUser.HomeDir)
	}
	return currentDir
}