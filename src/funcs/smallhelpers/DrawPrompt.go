package smallhelpers

import "CaSh/funcs/color"

func Drawprompt(user string, host string, currentDir string) {
	color.PrintBlue("[")
	color.PrintGreen(user)
	color.PrintBlue("@")
	color.PrintMagenta(host)
	color.PrintBlue("#")
	color.PrintYellow(currentDir)
	if user == "root" {
		color.PrintBlue("] >>> ")
	} else {
		color.PrintBlue("] >> ")
	}
}