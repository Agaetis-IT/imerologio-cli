package ime_generate

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Agaetis-IT/imerologio-cli/pkg/ime_types"
	"github.com/Agaetis-IT/imerologio-cli/pkg/ime_utils"
	"github.com/cheggaaa/pb/v3"
)

func GenerateSpringBoot(bar *pb.ProgressBar, answers ime_types.Answers) error {
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

func getStarter(bar *pb.ProgressBar, appPath string, answers ime_types.Answers) error {
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
	_, err = ime_utils.Unzip(starterZipPath, appPath)
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
