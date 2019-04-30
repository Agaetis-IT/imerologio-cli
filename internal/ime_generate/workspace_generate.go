package ime_generate

import (
	. "../../pkg/ime_types"
	. "../../pkg/ime_utils"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
)

func GenerateWorkspace(bar *pb.ProgressBar, path string) error {
	mkdirSucceeded := os.Mkdir(path, 0750)
	if mkdirSucceeded != nil {
		return mkdirSucceeded
	}
	bar.Increment()
	return nil
}

func RecapWorkspace(answers Answers) {
	PrintlnInfo("-- General")
	Print("App name: ")
	PrintlnPrompt(answers.Name)
	Print("App path: ")
	PrintlnPrompt(answers.Path)
}
