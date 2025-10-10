package color

import "fmt"

// I love ANSI Escape Codes

func PrintRed(printee string) {
	fmt.Print("\033[31m" + printee + "\033[0m")
}

func PrintBlue(printee string) {
	fmt.Print("\033[34m" + printee + "\033[0m")
}

func PrintGreen(printee string) {
	fmt.Print("\033[32m" + printee + "\033[0m")
}

func PrintYellow(printee string) {
	fmt.Print("\033[33m" + printee + "\033[0m")
}

func PrintMagenta(printee string) {
	fmt.Print("\033[35m" + printee + "\033[0m")
}

// Println variants that append a newline

func PrintRedln(printee string) {
	fmt.Println("\033[31m" + printee + "\033[0m")
}

func PrintBlueln(printee string) {
	fmt.Println("\033[34m" + printee + "\033[0m")
}

func PrintGreenln(printee string) {
	fmt.Println("\033[32m" + printee + "\033[0m")
}

func PrintYellowln(printee string) {
	fmt.Println("\033[33m" + printee + "\033[0m")
}

func PrintMagentaln(printee string) {
	fmt.Println("\033[35m" + printee + "\033[0m")
}
