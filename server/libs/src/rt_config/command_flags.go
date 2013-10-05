package rt_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CommandFlags struct {
	ConfigFile string
	SqlScripts string
}

// ConfigFileExists checks
func (cf *CommandFlags) ConfigFileExists() bool {
	return exists(cf.ConfigFile)
}

func (cf *CommandFlags) SqlScriptsExist() bool {
	return exists(cf.SqlScripts)
}

func (cf *CommandFlags) LoadConfig() *EnvironmentConfig {

	bytes, err := ioutil.ReadFile(cf.ConfigFile)

	if err != nil {
		fmt.Println(err)
	}

	val := &EnvironmentConfig{}
	err = json.Unmarshal(bytes, val)
	val.ConfigFile = cf.ConfigFile

	if err != nil {
		fmt.Println(err)
	}

	return val
}

func (cf *CommandFlags) CurrentConfiguration() *EnvironmentConfig {
	if config == nil {
		config = cf.LoadConfig()
	}
	return config
}

func (cf *CommandFlags) LoadEnvironmentConfig() *EnvironmentConfig {

	// Didn't find the config file path from the command line
	// so try the file by climbing the directory tree looking
	// for the config file.
	if !cf.ConfigFileExists() {
		cf.ConfigFile = findConfigFile(default_config_name)
	}

	if cf.ConfigFileExists() {
		return cf.CurrentConfiguration()
	}

	fmt.Println("Configuration doesn't exist: ", cf)

	return &EnvironmentConfig{}
}
