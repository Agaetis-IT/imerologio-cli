package ime_survey

import (
	"errors"
	"gopkg.in/AlecAivazis/survey.v1"
	"strings"
)

func AskLaunchGeneration(launchGeneration *bool) error {
	var answer = "n"
	var err = survey.AskOne(&survey.Input{Message: "Launch generation [y/n]"}, &answer,
		func(val interface{}) error {
			response := strings.ToLower(val.(string))
			if response != "n" && response != "no" && response != "y" && response != "yes" {
				return errors.New("only y, yes, n or no are allowed")
			}
			return nil
		})
	if err != nil {
		return err
	}
	*launchGeneration = strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes"
	return nil
}
