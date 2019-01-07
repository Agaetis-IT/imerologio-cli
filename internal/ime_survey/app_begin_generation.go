package ime_survey

import (
	"gopkg.in/AlecAivazis/survey.v1"
)

func AskBeginGeneration(launchGeneration *bool) error {
	var answer = false
	var err = survey.AskOne(&survey.Confirm{Message: "Begin generation?"}, &answer, nil)
	if err != nil {
		return err
	}
	*launchGeneration = answer
	return nil
}
