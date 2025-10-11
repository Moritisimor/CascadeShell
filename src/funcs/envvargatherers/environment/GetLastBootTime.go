package environment

import (
	"fmt"
	"log"
	"github.com/shirou/gopsutil/v4/host"
)

func GetLastBootTime() string {
	bootTime, err := host.BootTime()
	if err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%.2f Seconds", float64(bootTime) / 100000000)
}