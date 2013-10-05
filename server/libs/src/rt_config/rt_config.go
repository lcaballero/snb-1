package rt_config

import (
	"fmt"
	enc "json_helpers"
	"path"
	"path/filepath"
)

var (
	config *EnvironmentConfig = nil
)

type EnvironmentConfig struct {
	ConfigFile string
	Dev        RuntimeConfig
	Acc        RuntimeConfig
	Prod       RuntimeConfig
}

type RuntimeConfig struct {
	ConnectionString string
	SqlScripts       string
	WebServerPort    string
	DbServerPort     string
}

func Config() *EnvironmentConfig {

	if config != nil {
		return config
	}

	config = loadFromCommandLine()

	return config
}

func loadFromCommandLine() *EnvironmentConfig {

	cf := newFlags().createCommandFlags()

	configFileExists := cf.ConfigFileExists()

	// Didn't find the config file path from the command line
	// so try the file by climbing the directory tree looking
	// for the config file.
	if !configFileExists {
		cf.ConfigFile, configFileExists = findConfigFile(default_config_name)
	}

	if configFileExists {
		return cf.CurrentConfiguration()
	}

	fmt.Println("Configuration doesn't exist: ", cf)

	return &EnvironmentConfig{}
}

func findConfigFile(file string) (string, bool) {

	abs, _ := filepath.Abs(file)
	dir := path.Dir(abs)
	found := exists(abs)

	for !found {

		parent := dir + "/../"
		dir = path.Clean(parent)
		abs = path.Join(dir, file)

		found = exists(abs)

		if dir == "/" || dir == "" {
			break
		}
	}

	return abs, found
}

func (r *RuntimeConfig) String() string {
	return enc.ToIndentedJson(r, "", "   ")
}

func (env *EnvironmentConfig) String() string {
	js := enc.ToIndentedJson(env, "", "   ")
	return js
}
