package rt_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	enc "json_helpers"
)

type EnvironmentConfig struct {
	Dev, Acc, Prod RuntimeConfig
}

type RuntimeConfig struct {
	ConnectionString string
	SqlScripts       string
}

var config *EnvironmentConfig = nil

func init() {
	fmt.Println("rt_config.init()")

	config = loadConfig("../../../config.js")

	js := enc.ToIndentedJson(config, "", "   ")

	fmt.Println(js)
}

func loadConfig(path string) *EnvironmentConfig {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println(err)
	}

	val := &EnvironmentConfig{}
	err = json.Unmarshal(bytes, val)

	if err != nil {
		fmt.Println(err)
	}

	return val
}
