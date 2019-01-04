package ime_survey

import (
	. "../../pkg/ime_types"
	. "../../pkg/ime_utils"
	"errors"
	"gopkg.in/AlecAivazis/survey.v1"
	"os"
	"os/user"
	"path/filepath"
)

func AskAppPath(answers *Answers) error {
	questions := []*survey.Question{
		{
			Name:      "path",
			Prompt:    &survey.Input{Message: "App path [" + suggestAppPath(answers.Name) + "]:"},
			Validate:  validateAppPath(answers.Name),
			Transform: transformAppPath(answers.Name),
		},
	}
	return survey.Ask(questions, answers)
}

func suggestAppPath(appName string) string {
	usr, err := user.Current()
	if err != nil {
		PrintlnError("Error while retrieving user's home directory")
	}
	return filepath.Join(usr.HomeDir, appName)
}

func validateAppPath(appName string) func(interface{}) error {
	return func(appPathValue interface{}) error {
		appPath := transformAppPath(appName)(appPathValue).(string)

		// check that parent folder exists and that app folder does not exist to erase nothing
		parentPath := filepath.Dir(appPath)
		if _, err := os.Stat(parentPath); os.IsNotExist(err) {
			return errors.New("The parent folder must exist and " + parentPath + " does not")
		} else if _, err := os.Stat(appPath); !os.IsNotExist(err) {
			return errors.New("The given folder must not exist and " + appPath + " does")
		}

		return nil
	}
}

func transformAppPath(appName string) func(interface{}) interface{} {
	return func(appPathValue interface{}) interface{} {
		appPath := appPathValue.(string)
		if appPath == "" {
			return suggestAppPath(appName)
		}

		basePath := filepath.Base(appPath)
		// if user omitted the app name, let add it for him
		if basePath != appName {
			appPath = filepath.Join(appPath, appName)
		}

		return appPath
	}
}
