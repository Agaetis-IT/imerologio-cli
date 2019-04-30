package main

import (
	. "../../internal/ime_generate"
	. "../../internal/ime_survey"
	. "../../pkg/ime_types"
	. "../../pkg/ime_utils"
	"gopkg.in/cheggaaa/pb.v1"
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
		err = generateApp(answers)
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
	RecapWorkspace(answers)
	RecapKafka(answers)
	PrintlnInfo("")
}

func generateApp(answers Answers) error {
	count := 10
	bar := pb.StartNew(count)
	bar.ShowCounters = false

	err := GenerateWorkspace(bar, answers.Path)
	if err != nil {
		return err
	}

	err = GenerateKafka(bar, answers)
	if err != nil {
		return err
	}

	return nil
}
