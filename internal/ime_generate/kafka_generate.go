package ime_generate

import (
	. "../../pkg/ime_utils"
	"gopkg.in/cheggaaa/pb.v1"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const STRIMZI_VERSION = "0.11.1"
const STRIMZI_URL = "https://github.com/strimzi/strimzi-kafka-operator/releases/download/" + STRIMZI_VERSION + "/strimzi-" + STRIMZI_VERSION + ".zip"

func GenerateKafka(bar *pb.ProgressBar, workspacePath string, namespace string) error {
	strimziPath := workspacePath + "/event-store"
	err := getStrimzi(bar, workspacePath, strimziPath)
	if err != nil {
		return err
	}

	err = customizeStrimziClusterRole(bar, strimziPath+"/strimzi-"+STRIMZI_VERSION, namespace)
	if err != nil {
		return err
	}

	// TODO
	// customizeStrimziClusterRoleTopic
	// customizeStrimziClusterRoleUser
	// write shell that install Strimzi in a cluster

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
	bar.Increment()
	return err
}

func customizeStrimziClusterRole(bar *pb.ProgressBar, strimziPath string, namespace string) error {
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
			read, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			newContents := strings.Replace(string(read), "namespace: myproject", "namespace: "+namespace, -1)

			err = ioutil.WriteFile(path, []byte(newContents), 0)
			if err != nil {
				return nil
			}
		}
		bar.Increment()

		return nil
	}
}
