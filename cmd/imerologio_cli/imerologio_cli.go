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
	err := AskAppName(&answers)
	if err != nil {
		PrintlnError(err.Error())
		return
	}

	err = AskAppPath(&answers)
	if err != nil {
		PrintlnError(err.Error())
		return
	}

	err = AskKafkaNamespace(&answers)
	if err != nil {
		PrintlnError(err.Error())
		return
	}

	showRecap(answers)

	beginGeneration := false
	err = AskBeginGeneration(&beginGeneration)
	if err != nil {
		PrintlnError(err.Error())
		return
	}
	if !beginGeneration {
		PrintlnInfo("Ok, I've done nothing. See you soon 👋")
	} else {
		err = GenerateApp(answers)
		if err != nil {
			PrintlnError(err.Error())
			return
		}
		PrintlnInfo("Done! Happy coding 🎉")
	}
}

func showRecap(answers Answers) {
	PrintlnInfo("")
	PrintlnInfo("Ok, let's recap")
	PrintlnInfo("-- General")
	Print("App name: ")
	PrintlnPrompt(answers.Name)
	Print("App path: ")
	PrintlnPrompt(answers.Path)
	PrintlnInfo("-- Kafka")
	Print("Namespace: ")
	PrintlnPrompt(answers.KafkaNamespace)
	PrintlnInfo("")
}
