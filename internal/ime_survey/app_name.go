package ime_survey

import (
	. "../../pkg/ime_types"
	"gopkg.in/AlecAivazis/survey.v1"
)

func AskAppName(answers *Answers) error {
	return survey.AskOne(&survey.Input{Message: "App name:"}, &answers.Name, survey.Required)
}
