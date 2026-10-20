package storage

import (
	"cli/bins"
	"cli/file"
	"encoding/json"
	"errors"
	"os"
)

func BinToJson(jsonName string, bin *bins.Bin) error {
	if file.СheckJSONExtension(jsonName) {
		jsonBytes, err := json.Marshal(bin)
		if err != nil {
			return err
		}

		err = os.WriteFile(jsonName, jsonBytes, 0644)
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New("invalid extension")
}

func JsonToBin(jsonName string) (*bins.Bin, error) {
	if file.СheckJSONExtension(jsonName) {
		return bins.NewBinFromJson(jsonName)
	}
	return nil, errors.New("invalid extension")
}

func BinListToJson(jsonName string, binList *bins.BinList) error {
	if file.СheckJSONExtension(jsonName) {
		jsonBytes, err := json.Marshal(binList)
		if err != nil {
			return err
		}

		err = os.WriteFile(jsonName, jsonBytes, 0644)
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New("invalid extension")
}

func JsonToBinList(jsonName string) (*bins.BinList, error) {
	if file.СheckJSONExtension(jsonName) {
		return bins.NewBinListFromJson(jsonName)
	}
	return nil, errors.New("invalid extension")
}
