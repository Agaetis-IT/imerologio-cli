package ime_types

import (
	"github.com/AlecAivazis/survey/v2"
)

type Answers struct {
	Name                           string
	Path                           string
	KafkaOperatorNamespace         string
	KafkaClusterName               string
	KafkaClusterNamespace          string
	KafkaClusterPersistenceEnabled bool
	KafkaTopics                    string

	KafkaQuestions     []*survey.Question
	WorkspaceQuestions []*survey.Question
}
