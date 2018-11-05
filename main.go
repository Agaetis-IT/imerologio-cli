package main

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/fatih/color"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var promptColor = color.New(color.FgCyan).SprintfFunc()
var errorColor = color.New(color.FgRed).SprintfFunc()
var infoColor = color.New(color.FgMagenta).SprintfFunc()

func main() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.

	shell.Println(infoColor("Imerologio CLI helps you bootstrap your event sourcing app easily ✨"))

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "imerologio",
		Help: "Create a new imerologio application",
		Func: imerologio,
	})

	// when started with "exit" as first argument, assume non-interactive execution
	shell.Process("imerologio")
}

func imerologio(c *ishell.Context) {
	c.ShowPrompt(false)

	appName := getAppName(c)
	appPath := getAppPath(c, appName)

	c.Println("")
	c.Println(infoColor("Ok, let's recap"))
	c.Println("----------------")
	c.Print(promptColor("App name : "))
	c.Println(appName)
	c.Print(promptColor("App path : "))
	c.Println(appPath)
	c.Println("----------------")
	for {
		c.Print(infoColor("Launch generation ? [y/n]: "))
		goGeneration := c.ReadLine()
		if strings.ToLower(goGeneration) == "n" || strings.ToLower(goGeneration) == "no" {
			c.Print("Ok, I've done nothing. See you soon 👋")
			break
		} else if strings.ToLower(goGeneration) == "y" || strings.ToLower(goGeneration) == "yes" {
			c.ProgressBar().Start()
			for i := 0; i < 101; i++ {
				c.ProgressBar().Suffix(fmt.Sprint(" ", i, "%"))
				c.ProgressBar().Progress(i)
				generateApp(appName, appPath)
			}
			c.ProgressBar().Stop()
			c.Println("")
			c.Print("Happy coding 🎉")
			break
		}
		c.Println("Only y, yes, n or no are allowed")
	}
}

func getAppName(c *ishell.Context) string {
	for {
		c.Print(promptColor("App name: "))
		result := c.ReadLine()
		if result != "" {
			return result
		}
		c.Println(errorColor("App name must be a non empty string"))
	}
}

func getAppPath(c *ishell.Context, appName string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Error while retrieving user's home directory")
	}
	suggestedPath := filepath.Join(usr.HomeDir, appName)

	for {
		c.Print(promptColor("Path [%s]: ", suggestedPath))
		var appPath = c.ReadLine()

		// user chose the suggested path
		if appPath == "" {
			return suggestedPath
		}

		basePath := filepath.Base(appPath)
		// if user omitted the app name, let add it for him
		if basePath != appName {
			appPath = filepath.Join(appPath, appName)
		}

		// check that parent folder exists and that app folder does not exist to erase nothing
		parentPath := filepath.Dir(appPath)
		if _, err := os.Stat(parentPath); os.IsNotExist(err) {
			c.Println(errorColor("The parent folder must exist and %s does not", parentPath))
		} else if _, err := os.Stat(appPath); !os.IsNotExist(err) {
			c.Println(errorColor("The given folder must be empty and %s is not", appPath))
		} else {
			return appPath
		}
	}
}

func generateApp(appName string, appPath string) {
	os.Mkdir(appPath, os.ModeDir)
}
