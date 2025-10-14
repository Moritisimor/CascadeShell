package smallhelpers

import (
	"CaSh/funcs/color"
)

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

func SPrintPrompt(user string, host string, currentDir string) string {
	end := color.SprintBlue("] >> ")
	if user == "root" {
		end = color.SprintBlue("] >>> ")
	}
	return 	color.SprintBlue("[") 			+
			color.SprintGreen(user) 		+
			color.SprintBlue("@")			+
			color.SprintMagenta(host)		+
			color.SprintBlue("#")			+
			color.SprintYellow(currentDir)	+
			end
}
