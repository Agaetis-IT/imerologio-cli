package ime_generate

import (
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/Agaetis-IT/imerologio-cli/pkg/ime_types"
	"github.com/AlecAivazis/survey/v2"
	"github.com/cheggaaa/pb/v3"
)

type manifestReader struct {
	answers        *ime_types.Answers
	strimziVersion string
}

func dumpQuestionParameters(answers ime_types.Answers, questions []*survey.Question, result *string) {
	value := reflect.ValueOf(answers)

	for _, q := range questions {
		field := value.FieldByName(q.Name)
		switch field.Type().String() {
		case "bool":
			*result += q.Name + ": " + strconv.FormatBool(field.Bool()) + "\n"
		default:
			*result += q.Name + ": " + field.String() + "\n"
		}
	}
}

func dumpToFile(manifest manifestReader) []byte {
	result := "Name: " + manifest.answers.Name + "\n"

	//Dump Workspace info
	// go over every question
	dumpQuestionParameters(*manifest.answers, manifest.answers.WorkspaceQuestions, &result)

	//Dump Kafka information
	dumpQuestionParameters(*manifest.answers, manifest.answers.KafkaQuestions, &result)

	result += "Strimzi Version: " + STRIMZI_VERSION + "\n"
	result += "Strimzi Url: " + STRIMZI_URL + "\n"
	return []byte(result)

}

// GenerateManifest file to dump the elements used and created
func GenerateManifest(bar *pb.ProgressBar, answers *ime_types.Answers) error {
	var manifestReader manifestReader
	manifestReader.answers = answers
	manifestReader.strimziVersion = STRIMZI_VERSION
	manifestFilePath := answers.Path + "/" + answers.Name + ".manifest"

	// Write the body to file
	err := ioutil.WriteFile(manifestFilePath, dumpToFile(manifestReader), 0644)
	if err != nil {
		return err
	}

	bar.Increment()
	return nil
}
