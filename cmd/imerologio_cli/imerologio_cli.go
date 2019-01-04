package main

import (
	. "../../internal/ime_generate"
	. "../../internal/ime_survey"
	. "../../pkg/ime_types"
	. "../../pkg/ime_utils"
)

func main() {
	PrintlnInfo("Imerologio CLI helps you bootstrap your event sourcing app easily ✨")

	answers := Answers{}
	//
	var err = AskAppName(&answers)
	if err != nil {
		PrintlnError(err.Error())
		return
	}

	err = AskAppPath(&answers)
	if err != nil {
		PrintlnError(err.Error())
		return
	}

	showRecap(answers)

	var launchGeneration = false
	err = AskLaunchGeneration(&launchGeneration)
	if err != nil {
		PrintlnError(err.Error())
		return
	}
	if !launchGeneration {
		PrintlnInfo("Ok, I've done nothing. See you soon 👋")
	} else {
		err = GenerateApp(answers)
		if err != nil {
			PrintlnError(err.Error())
			return
		}
		PrintlnInfo("Done ! Happy coding 🎉")
	}
}

func showRecap(answers Answers) {
	PrintlnInfo("----------------")
	PrintlnInfo("Ok, let's recap")
	Print("App name: ")
	PrintlnPrompt(answers.Name)
	Print("App path: ")
	PrintlnPrompt(answers.Path)
	PrintlnInfo("----------------")
}
