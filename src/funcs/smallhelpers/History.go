package smallhelpers

import (
	"CaSh/funcs/color"
	"CaSh/funcs/envvargatherers/environment"
	"bufio"
	"fmt"
	"log"
	"os"
)

func MakeHistory() {
	_, dirErr := os.Stat(environment.GetUser().HomeDir + string(os.PathSeparator) + ".cash")
	if os.IsNotExist(dirErr) {
		directoryErr := os.Mkdir(environment.GetUser().HomeDir + string(os.PathSeparator) + ".cash", 0755)
		if directoryErr != nil {
			log.Fatalf("Could not create .cash directory!\nError: %s", directoryErr.Error())
		}
	}

	_, fileErr := os.Stat(environment.GetUser().HomeDir + string(os.PathSeparator) + ".cash" + string(os.PathSeparator) + "cashhistory")
	if os.IsNotExist(fileErr) {
		_, creationErr := os.Create(environment.GetUser().HomeDir + string(os.PathSeparator) + ".cash" + string(os.PathSeparator) + "cashhistory")
		if creationErr != nil {
			log.Fatalf("Could not create cashhistory file!\nError: %s", creationErr.Error())
		}
	}
}

func GetHistory() string {
	return environment.GetUser().HomeDir + string(os.PathSeparator) + ".cash" + string(os.PathSeparator) + "cashhistory"
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
	var response string
	color.PrintBlue("Clear History?\nY/N: ")
	fmt.Scanf("%s", &response)

	if Normalize(response) == "yes" || Normalize(response) == "y" {
		err := os.Remove(GetHistory())
		if err != nil {
			log.Fatal(err.Error())
		}

		color.PrintGreenln("Successfully cleared history!")
		return
	}

	color.PrintBlueln("Cancelled history clear!")
}
