package helpers

import (
	"github.com/Moritisimor/EpsilonFetch/pkg/color"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/user"
)

func MakeHistory() {
	currentUser, _ := user.Current()
	_, dirErr := os.Stat(currentUser.HomeDir + string(os.PathSeparator) + ".cash")
	if os.IsNotExist(dirErr) {
		directoryErr := os.Mkdir(currentUser.HomeDir + string(os.PathSeparator) + ".cash", 0755)
		if directoryErr != nil {
			log.Fatalf("Could not create .cash directory!\nError: %s", directoryErr.Error())
		}
	}

	_, fileErr := os.Stat(currentUser.HomeDir + string(os.PathSeparator) + ".cash" + string(os.PathSeparator) + "cashhistory")
	if os.IsNotExist(fileErr) {
		_, creationErr := os.Create(currentUser.HomeDir + string(os.PathSeparator) + ".cash" + string(os.PathSeparator) + "cashhistory")
		if creationErr != nil {
			log.Fatalf("Could not create cashhistory file!\nError: %s", creationErr.Error())
		}
	}
}

func GetHistory() string {
	currentUser, _ := user.Current()
	return currentUser.HomeDir + string(os.PathSeparator) + ".cash" + string(os.PathSeparator) + "cashhistory"
}

func ReadHistory() {
	openedFile, openErr := os.Open(GetHistory())
	if openErr != nil {
		log.Fatal(openErr.Error())
	}
	defer openedFile.Close()

	counter := 1
	scanner := bufio.NewScanner(openedFile)
	for scanner.Scan() {
		color.PrintBlueln(fmt.Sprintf("%d: %s", counter, string(scanner.Text())))
		counter++
	}
}

func ClearHistory() {
	err := os.Remove(GetHistory())
	if err != nil {
		log.Fatal(err.Error())
	}

	color.PrintGreenln("Successfully cleared history!")
}
