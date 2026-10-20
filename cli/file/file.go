package file

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ReadSomethingFile(nameOfFile string) ([]byte, error) {
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

func СheckJSONExtension(nameOfFile string) bool {
	ext := filepath.Ext(nameOfFile)
	return strings.EqualFold(ext, ".json")
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
