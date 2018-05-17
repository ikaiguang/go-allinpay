package config

import (
	"runtime"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

const jsonFile = "conf.json"

// set app run path, notice : runtime.Caller param skip
func init() {
	// get path
	_, file, _, _ := runtime.Caller(0)
	currentPath := filepath.Join(file, ".."+string(filepath.Separator))
	jsonFile := filepath.Join(currentPath, jsonFile)
	// read json
	jsonByte, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Println("ioutil.ReadFile error : ", err.Error())
	}
	// decode json
	cfg := make(map[string]string)
	if err := json.Unmarshal(jsonByte, &cfg); err != nil {
		fmt.Println("json.Unmarshal error : ", err.Error())
	}
	// set env
	for key, value := range cfg {
		if err := os.Setenv(key, value); err != nil {
			fmt.Println("os.Setenv error : ", err.Error())
		}
	}
}
