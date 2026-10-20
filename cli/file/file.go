package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func readSomethingFile(nameOfFile string) ([]byte, error) {
	if ok, err := checkFile(nameOfFile); ok {
		somthingBytes, err := os.ReadFile(nameOfFile)
		if err != nil {
			return nil, err
		}
		return somthingBytes, nil
	} else {
		return nil, err
	}

}

func checkJSONExtension(nameOfFile string) (bool, error) {
	if ok, err := checkFile(nameOfFile); ok {
		ext := filepath.Ext(nameOfFile)
		return strings.EqualFold(ext, ".json"), nil
	} else {
		return false, err
	}
}

func checkFile(nameOfFile string) (bool, error) {
	info, err := os.Stat(nameOfFile)
	if err != nil {
		return false, err
	} else if info.IsDir() {
		return false, errors.New("no such files, it is directory")
	}

	return true, nil
}
