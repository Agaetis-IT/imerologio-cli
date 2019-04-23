package ime_generate

import (
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
