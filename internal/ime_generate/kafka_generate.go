package ime_generate

import (
	. "../../pkg/ime_utils"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"net/http"
	"os"
)

const strimziUrl = "https://github.com/strimzi/strimzi-kafka-operator/releases/download/0.11.1/strimzi-0.11.1.zip"

func GenerateKafka(bar *pb.ProgressBar, workspacePath string) error {
	err := getStrimzi(bar, workspacePath)
	if err != nil {
		return err
	}
	return nil
}

func getStrimzi(bar *pb.ProgressBar, workspace string) error {
	strimziZipPath := workspace + "/strimzi.zip"
	strimziPath := workspace + "/strimzi"

	// Download strimzi release
	resp, err := http.Get(strimziUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bar.Increment()

	// Create the file
	out, err := os.Create(strimziZipPath)
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
	_, err = Unzip(strimziZipPath, strimziPath)
	bar.Increment()
	return err
}
