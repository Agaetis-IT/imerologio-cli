package main

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/cheggaaa/pb.v1"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

var promptColor = color.New(color.FgCyan).SprintfFunc()
var infoColor = color.New(color.FgMagenta).SprintfFunc()

type Answers struct {
	Name string
	Path string
}

func main() {
	fmt.Println(infoColor("Imerologio CLI helps you bootstrap your event sourcing app easily ✨"))

	answers := Answers{}

	var err = survey.AskOne(&survey.Input{Message: "App name:"}, &answers.Name, survey.Required)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	questions := []*survey.Question{
		{
			Name:      "path",
			Prompt:    &survey.Input{Message: "App path [" + suggestAppPath(answers.Name) + "]:"},
			Validate:  validateAppPath(answers.Name),
			Transform: transformAppPath(answers.Name),
		},
	}
	err = survey.Ask(questions, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	showRecap(answers)

	var launchGeneration = "n"
	err = survey.AskOne(&survey.Input{Message: "Launch generation [y/n]"}, &launchGeneration, func(val interface{}) error {
		response := strings.ToLower(val.(string))
		if response != "n" && response != "no" && response != "y" && response != "yes" {
			return errors.New("Only y, yes, n or no are allowed")
		}
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if strings.ToLower(launchGeneration) == "n" || strings.ToLower(launchGeneration) == "no" {
		fmt.Println("Ok, I've done nothing. See you soon 👋")
	} else {
		count := 5000
		bar := pb.StartNew(count)
		bar.ShowCounters = false
		for i := 0; i < count; i++ {
			bar.Increment()
			time.Sleep(time.Millisecond)
		}
		generateApp(answers)
		bar.FinishPrint(infoColor("Happy coding 🎉"))
	}
}

func suggestAppPath(appName string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Error while retrieving user's home directory")
	}
	return filepath.Join(usr.HomeDir, appName)
}

func validateAppPath(appName string) func(interface{}) error {
	return func(appPathValue interface{}) error {
		appPath := appPathValue.(string)
		if appPath == "" {
			return nil
		}

		basePath := filepath.Base(appPath)
		// if user omitted the app name, let add it for him
		if basePath != appName {
			appPath = filepath.Join(appPath, appName)
		}

		// check that parent folder exists and that app folder does not exist to erase nothing
		parentPath := filepath.Dir(appPath)
		if _, err := os.Stat(parentPath); os.IsNotExist(err) {
			return errors.New("The parent folder must exist and " + parentPath + " does not")
		} else if _, err := os.Stat(appPath); os.IsExist(err) {
			return errors.New("The given folder must be empty and " + appPath + " is not")
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

func showRecap(answers Answers) {
	fmt.Println("")
	fmt.Println(infoColor("Ok, let's recap"))
	fmt.Println("----------------")
	fmt.Print(promptColor("App name: "))
	fmt.Println(answers.Name)
	fmt.Print(promptColor("App path: "))
	fmt.Println(answers.Path)
	fmt.Println("----------------")
}

func generateApp(answers Answers) {
	os.Mkdir(answers.Path, 0700)
}
