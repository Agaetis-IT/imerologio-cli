package ime_utils

import (
	"io"
	"io/ioutil"
	"os"
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

func CopyFile(source string, destination string) error {
	from, err := os.Open(source)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		return err
	}

	return nil
}
