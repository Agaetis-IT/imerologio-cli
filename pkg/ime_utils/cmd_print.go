package ime_utils

import (
	"fmt"
	"github.com/fatih/color"
)

var promptColor = color.New(color.FgCyan).SprintfFunc()
var errorColor = color.New(color.FgRed).SprintfFunc()
var infoColor = color.New(color.FgMagenta).SprintfFunc()

func Print(message string) {
	fmt.Print(message)
}

func PrintlnPrompt(message string) {
	fmt.Println(promptColor(message))
}

func PrintlnError(message string) {
	fmt.Println(errorColor(message))
}

func PrintlnInfo(message string) {
	fmt.Println(infoColor(message))
}
