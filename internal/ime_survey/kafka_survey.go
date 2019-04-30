package ime_survey

import (
	. "../../pkg/ime_types"
	"gopkg.in/AlecAivazis/survey.v1"
)

func AskKafkaNamespace(answers *Answers) error {
	questions := []*survey.Question{
		{
			Name:      "KafkaOperatorNamespace",
			Prompt:    &survey.Input{Message: "Kafka Operator - namespace [default]:"},
			Transform: transformDefaultValue("default"),
		},
		{
			Name:      "KafkaClusterName",
			Prompt:    &survey.Input{Message: "Kafka Cluster - name [cluster-" + answers.Name + "]:"},
			Transform: transformDefaultValue("cluster-" + answers.Name),
		},
		{
			Name:      "KafkaClusterNamespace",
			Prompt:    &survey.Input{Message: "Kafka Cluster - namespace [default]:"},
			Transform: transformDefaultValue("default"),
		},
		{
			Name:   "KafkaClusterPersistenceEnabled",
			Prompt: &survey.Confirm{Message: "Kafka Cluster - enable persistence ?"},
		},
	}
	return survey.Ask(questions, answers)
}

func transformDefaultValue(defaultValue string) func(interface{}) interface{} {
	return func(value interface{}) interface{} {
		valueAsString := value.(string)
		if valueAsString == "" {
			return defaultValue
		}

		return valueAsString
	}
}
