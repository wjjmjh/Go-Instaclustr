package util

import (
	"Go-Instaclustr/src"
	"encoding/json"
	"io/ioutil"
	"os"
)

func OpenConfigJson(fn string) (*src.Configuration, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var config src.Configuration
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &config)
	return &config, nil
}
