package color

import (
	"fmt"
	"runtime"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

func RedText(text string) string {
	return fmt.Sprintf("%s%s%s", Red, text, Reset)
}

func GreenText(text string) string {
	return fmt.Sprintf("%s%s%s", Green, text, Reset)
}

func YellowText(text string) string {
	return fmt.Sprintf("%s%s%s", Yellow, text, Reset)
}

func BlueText(text string) string {
	return fmt.Sprintf("%s%s%s", Blue, text, Reset)
}

func CyanText(text string) string {
	return fmt.Sprintf("%s%s%s", Cyan, text, Reset)
}
