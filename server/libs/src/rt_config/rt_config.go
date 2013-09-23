package rt_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	enc "json_helpers"
	"path"
)

type EnvironmentConfig struct {
	Dev, Acc, Prod RuntimeConfig
}

type RuntimeConfig struct {
	ConnectionString string
	SqlScripts       string
	WebServerPort    string
	DbServerPort     string
}

var config *EnvironmentConfig = nil
var default_config_file string = "config.js"

func init() {

	cfg_file := path.Join("../../../", default_config_file)

	fmt.Println("Loading configuration file: ", cfg_file)

	config = loadConfig(cfg_file)
}

func DumpConfigFile() {

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
