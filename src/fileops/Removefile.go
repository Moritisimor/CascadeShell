package fileops

import (
	"os"
	"fmt"
)

func Removefile(target string) {
	err := os.Remove(target); if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Successfully deleted %s.\n", target)
}
