package ime_generate

import (
	. "../../pkg/ime_types"
	. "../../pkg/ime_utils"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const STRIMZI_VERSION = "0.11.1"
const STRIMZI_URL = "https://github.com/strimzi/strimzi-kafka-operator/releases/download/" + STRIMZI_VERSION + "/strimzi-" + STRIMZI_VERSION + ".zip"

func GenerateKafka(bar *pb.ProgressBar, answers Answers) error {
	workspacePath := answers.Path
	eventStorePath := workspacePath + "/event-store"
	strimziPath := eventStorePath + "/strimzi-" + STRIMZI_VERSION
	err := getStrimzi(bar, workspacePath, eventStorePath)
	if err != nil {
		return err
	}

	err = customizeClusterRole(bar, strimziPath, answers.KafkaOperatorNamespace)
	if err != nil {
		return err
	}

	err = customizeCluster(bar, strimziPath, answers.KafkaClusterName)
	if err != nil {
		return err
	}

	topics := getTopics(answers.KafkaTopics)
	for _, topic := range topics {
		err = customizeTopic(bar, strimziPath, topic, answers.KafkaClusterName)
		if err != nil {
			return err
		}
	}

	err = initializeDeploymentScript(bar, eventStorePath, strimziPath, answers)
	if err != nil {
		return err
	}

	return nil
}

func getStrimzi(bar *pb.ProgressBar, workspacePath string, strimziPath string) error {
	strimziZipPath := workspacePath + "/strimzi.zip"

	// Download strimzi release
	resp, err := http.Get(STRIMZI_URL)
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
	if err != nil {
		return err
	}
	bar.Increment()

	// Remove zip archive
	err = os.Remove(strimziZipPath)
	if err != nil {
		return err
	}
	bar.Increment()

	return nil
}

func customizeClusterRole(bar *pb.ProgressBar, strimziPath string, namespace string) error {
	err := filepath.Walk(strimziPath+"/install/cluster-operator", replaceNamespaceInRoleBindings(namespace, bar))
	if err != nil {
		return err
	}
	bar.Increment()
	return nil
}

func replaceNamespaceInRoleBindings(namespace string, bar *pb.ProgressBar) func(string, os.FileInfo, error) error {
	return func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !!fi.IsDir() {
			return nil
		}

		matched, err := filepath.Match("*RoleBinding*.yaml", fi.Name())

		if err != nil {
			return err
		}

		if matched {
			err = ReplaceInFile(path, "namespace: myproject", "namespace: "+namespace)
			if err != nil {
				return nil
			}
		}
		bar.Increment()

		return nil
	}
}

func customizeCluster(bar *pb.ProgressBar, strimziPath string, clusterName string) error {
	err := filepath.Walk(strimziPath+"/examples/kafka", replaceClusterNameInKafkaExamples(clusterName, bar))

	if err != nil {
		return err
	}
	bar.Increment()
	return nil
}

func replaceClusterNameInKafkaExamples(clusterName string, bar *pb.ProgressBar) func(string, os.FileInfo, error) error {
	return func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !!fi.IsDir() {
			return nil
		}

		matched, err := filepath.Match("*.yaml", fi.Name())

		if err != nil {
			return err
		}

		if matched {
			err = ReplaceInFile(path, "my-cluster", clusterName)
			if err != nil {
				return nil
			}
		}
		bar.Increment()

		return nil
	}
}

func customizeTopic(bar *pb.ProgressBar, strimziPath string, topic string, clusterName string) error {
	err := CopyFile(strimziPath+"/examples/topic/kafka-topic.yaml", strimziPath+"/examples/topic/kafka-topic-"+topic+".yaml")
	if err != nil {
		return err
	}
	bar.Increment()

	err = ReplaceInFile(strimziPath+"/examples/topic/kafka-topic-"+topic+".yaml", "my-topic", topic)
	if err != nil {
		return err
	}
	bar.Increment()

	err = ReplaceInFile(strimziPath+"/examples/topic/kafka-topic-"+topic+".yaml", "my-cluster", clusterName)
	if err != nil {
		return err
	}
	bar.Increment()

	return nil
}

func initializeDeploymentScript(bar *pb.ProgressBar, eventStorePath string, strimziPath string, answers Answers) error {
	scriptName := eventStorePath + "/deploy_event_store.sh"
	script := "#!/bin/bash\n\n"

	script += "# Install cluster operator to expose Kafka cluster resources\n"
	script += "kubectl apply -f " + strimziPath + "/install/cluster-operator -n " + answers.KafkaOperatorNamespace + "\n\n"

	if answers.KafkaClusterPersistenceEnabled {
		script += "# Apply cluster with persistence\n"
		script += "kubectl apply -f " + strimziPath + "/examples/kafka/kafka-persistent.yaml -n " + answers.KafkaClusterNamespace + "\n\n"
	} else {
		script += "# Apply cluster without persistence\n"
		script += "kubectl apply -f " + strimziPath + "/examples/kafka/kafka-ephemeral.yaml -n " + answers.KafkaClusterNamespace + "\n\n"
	}

	topics := getTopics(answers.KafkaTopics)
	if len(topics) > 0 {
		script += "# Apply topics\n"
	}
	for _, topic := range topics {
		script += "kubectl apply -f " + strimziPath + "/examples/topic/kafka-topic-" + topic + ".yaml\n"
	}

	err := ioutil.WriteFile(scriptName, []byte(script), 0700)
	if err != nil {
		return err
	}
	bar.Increment()

	return nil
}

func RecapKafka(answers Answers) {
	PrintlnInfo("-- Kafka")
	PrintlnInfo("--- Operator")
	Print("Namespace: ")
	PrintlnPrompt(answers.KafkaOperatorNamespace)
	PrintlnInfo("--- Cluster")
	Print("Name: ")
	PrintlnPrompt(answers.KafkaClusterName)
	Print("Namespace: ")
	PrintlnPrompt(answers.KafkaClusterNamespace)
	Print("Persistence: ")
	PrintlnPrompt(strconv.FormatBool(answers.KafkaClusterPersistenceEnabled))
	topics := getTopics(answers.KafkaTopics)
	if len(topics) > 0 {
		Println("Topics: ")
		for _, topic := range topics {
			PrintlnPrompt("  - " + topic)
		}
	} else {
		Println("Topics: no topics")
	}
}

func getTopics(topicsAsString string) []string {
	return Split(topicsAsString, ',')
}
