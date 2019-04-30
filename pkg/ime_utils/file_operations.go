package ime_utils

import (
	"io/ioutil"
	"strings"
)

func ReplaceInFile(path string, old string, new string) error {
	read, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	newContents := strings.Replace(string(read), old, new, -1)

	return ioutil.WriteFile(path, []byte(newContents), 0)
}
