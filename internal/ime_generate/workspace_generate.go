package ime_generate

import (
	"os"

	"github.com/Agaetis-IT/imerologio-cli/pkg/ime_types"
	"github.com/Agaetis-IT/imerologio-cli/pkg/ime_utils"
	"github.com/cheggaaa/pb/v3"
)

func GenerateWorkspace(bar *pb.ProgressBar, path string) error {
	mkdirSucceeded := os.Mkdir(path, 0750)
	if mkdirSucceeded != nil {
		return mkdirSucceeded
	}
	bar.Increment()
	return nil
}

func RecapWorkspace(answers ime_types.Answers) {
	ime_utils.PrintlnInfo("-- General")
	ime_utils.Print("App name: ")
	ime_utils.PrintlnPrompt(answers.Name)
	ime_utils.Print("App path: ")
	ime_utils.PrintlnPrompt(answers.Path)
}
