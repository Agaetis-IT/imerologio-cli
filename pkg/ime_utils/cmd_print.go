package ime_utils

import (
	"fmt"
	"github.com/fatih/color"
)

func Print(message string) {
	fmt.Print(message)
}

func PrintlnPrompt(message string) {
	promptColor := color.New(color.FgCyan).SprintfFunc()
	fmt.Println(promptColor(message))
}

func PrintlnError(message string) {
	errorColor := color.New(color.FgRed).SprintfFunc()
	fmt.Println(errorColor(message))
}

func PrintlnInfo(message string) {
	infoColor := color.New(color.FgMagenta).SprintfFunc()
	fmt.Println(infoColor(message))
}
