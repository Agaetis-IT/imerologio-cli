package ime_survey

import (
	. "../../pkg/ime_types"
	"gopkg.in/AlecAivazis/survey.v1"
)

func AskKafkaNamespace(answers *Answers) error {
	questions := []*survey.Question{
		{
			Name:      "KafkaNamespace",
			Prompt:    &survey.Input{Message: "Kafka namespace [default]:"},
			Transform: transformKafkaNamespace(),
		},
	}
	return survey.Ask(questions, answers)
}

func transformKafkaNamespace() func(interface{}) interface{} {
	return func(kafkaNamespaceValue interface{}) interface{} {
		kafkaNamespace := kafkaNamespaceValue.(string)
		if kafkaNamespace == "" {
			return "default"
		}

		return kafkaNamespace
	}
}
