package ime_generate

import (
	"io"
	"net/http"
	"os"
	"strings"

	. "../../pkg/ime_types"
	. "../../pkg/ime_utils"
	"gopkg.in/cheggaaa/pb.v1"
)

func GenerateSpringBoot(bar *pb.ProgressBar, answers Answers) error {
	workspacePath := answers.Path
	appPath := workspacePath + "/app"

	err := os.Mkdir(appPath, 0750)
	if err != nil {
		return err
	}
	bar.Increment()

	err = getStarter(bar, appPath, answers)
	if err != nil {
		return err
	}

	return nil
}

func getStarter(bar *pb.ProgressBar, appPath string, answers Answers) error {
	starterZipPath := appPath + "/starter.zip"

	body := strings.NewReader(`dependencies=web,actuator&language=java&type=maven-project&baseDir=` + answers.Name)
	req, err := http.NewRequest("POST", "https://start.spring.io/starter.zip", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bar.Increment()

	// Create the file
	out, err := os.Create(starterZipPath)
	if err != nil {
		return err
	}
	defer out.Close()
	bar.Increment()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	bar.Increment()

	// Extract zip archive
	_, err = Unzip(starterZipPath, appPath)
	if err != nil {
		return err
	}
	bar.Increment()

	// Remove zip archive
	err = os.Remove(starterZipPath)
	if err != nil {
		return err
	}
	bar.Increment()

	return nil
}
