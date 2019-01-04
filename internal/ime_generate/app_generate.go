package ime_generate

import (
	. "../../pkg/ime_types"
	"gopkg.in/cheggaaa/pb.v1"
	"os"
	"time"
)

func GenerateApp(answers Answers) error {
	count := 2000
	bar := pb.StartNew(count)
	bar.ShowCounters = false
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	return os.Mkdir(answers.Path, 0750)
}
