package main

import (
	"github.com/Agaetis-IT/imerologio-cli/internal/ime_generate"
	"github.com/Agaetis-IT/imerologio-cli/internal/ime_survey"
	"github.com/Agaetis-IT/imerologio-cli/pkg/ime_types"
	"github.com/Agaetis-IT/imerologio-cli/pkg/ime_utils"
	"github.com/cheggaaa/pb/v3"
)

func main() {
	ime_utils.PrintlnInfo("Imerologio CLI helps you bootstrap your event sourcing app easily ✨")

	answers := ime_types.Answers{}
	err := ime_survey.AskAppName(&answers)
	if err != nil {
		ime_utils.PrintlnError(err.Error())
		return
	}

	err = ime_survey.AskAppPath(&answers)
	if err != nil {
		ime_utils.PrintlnError(err.Error())
		return
	}

	err = ime_survey.AskKafka(&answers)
	if err != nil {
		ime_utils.PrintlnError(err.Error())
		return
	}

	showRecap(answers)

	beginGeneration := false
	err = ime_survey.AskBeginGeneration(&beginGeneration)
	if err != nil {
		ime_utils.PrintlnError(err.Error())
		return
	}
	if !beginGeneration {
		ime_utils.PrintlnInfo("Ok, I've done nothing. See you soon 👋")
	} else {
		err = generateApp(answers)
		if err != nil {
			ime_utils.PrintlnError(err.Error())
			return
		}
		ime_utils.PrintlnInfo("Done! Happy coding 🎉")
	}
}

func showRecap(answers ime_types.Answers) {
	ime_utils.PrintlnInfo("")
	ime_utils.PrintlnInfo("Ok, let's recap")
	ime_generate.RecapWorkspace(answers)
	ime_generate.RecapKafka(answers)
	ime_utils.PrintlnInfo("")
}

func generateApp(answers ime_types.Answers) error {
	count := 11
	bar := pb.Simple.Start(count)
	bar.SetTemplateString(`{{string . "prefix"}}{{bar . }} {{percent . }}{{string . "suffix"}}`)

	err := ime_generate.GenerateWorkspace(bar, answers.Path)
	if err != nil {
		return err
	}

	err = ime_generate.GenerateKafka(bar, answers)
	if err != nil {
		return err
	}

	err = ime_generate.GenerateManifest(bar, answers)
	if err != nil {
		return err
	}

	return nil
}
