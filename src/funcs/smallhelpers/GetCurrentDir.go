package smallhelpers

import "os"

func GetCurrentDir() string {
	currentDir, _ := os.Getwd()
	return currentDir
}