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
