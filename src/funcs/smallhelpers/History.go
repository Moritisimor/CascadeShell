package smallhelpers

import (
	"CaSh/funcs/envvargatherers/environment"
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
