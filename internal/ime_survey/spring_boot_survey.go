package ime_survey

import (
	. "../../pkg/ime_types"
	"gopkg.in/AlecAivazis/survey.v1"
)

func AskSpringBoot(answers *Answers) error {
	questions := []*survey.Question{}
	return survey.Ask(questions, answers)
}
