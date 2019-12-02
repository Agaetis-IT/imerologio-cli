package ime_survey

import (
	"github.com/Agaetis-IT/imerologio-cli/pkg/ime_types"
	"github.com/AlecAivazis/survey/v2"
)

func AskSpringBoot(answers *ime_types.Answers) error {
	questions := []*survey.Question{}
	return survey.Ask(questions, answers)
}
