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

type ConfigPathProvider func() string

var config *EnvironmentConfig = nil
var default_config_name string = "config.js"
var default_config_root string = "../../../"

var PathProvider ConfigPathProvider = defaultPathProvider

func CurrentConfiguration() *EnvironmentConfig {
	if config == nil {
		config = LoadConfig(PathProvider())
	}
	return config
}

func defaultPathProvider() string {
	return path.Join(default_config_root, default_config_name)
}

func DumpConfigFile() {

	cf := CurrentConfiguration()
	js := enc.ToIndentedJson(cf, "", "   ")

	fmt.Println(js)
}

func LoadConfig(path string) *EnvironmentConfig {
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
