package storage

import (
	"cli/bins"
	"encoding/json"
	"os"
)

func BinToJson(jsonName string, bin *bins.Bin) error {
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

func JsonToBin(jsonName string) (*bins.Bin, error) {
	return bins.NewBinFromJson(jsonName)
}

func BinListToJson(jsonName string, binList *bins.BinList) error {
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

func JsonToBinList(jsonName string) (*bins.BinList, error) {
	return bins.NewBinListFromJson(jsonName)
}
