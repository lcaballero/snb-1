package rt_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	enc "json_helpers"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	config *EnvironmentConfig = nil
)

const (
	default_config_name = "config.js"
	default_config_root = "./"
	flag_config_file    = "--config-file"
	flag_sql_scripts    = "--sql-scripts"
)

type flags struct {
	configFile string
	sqlScripts string
}

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

		if dir == "/" {
			break
		}
	}

	return abs, found
}

func (r *RuntimeConfig) String() string {
	return enc.ToIndentedJson(r, "", "   ")
}

func newFlags() *flags {
	return &flags{
		configFile: flag_config_file,
		sqlScripts: flag_sql_scripts,
	}
}

func (f *flags) createCommandFlags() *CommandFlags {

	cf := &CommandFlags{}

	cf.ConfigFile, _ = FindFlag(f.configFile, os.Args)
	cf.SqlScripts, _ = FindFlag(f.sqlScripts, os.Args)

	return cf
}

func FindFlag(flag string, args []string) (string, bool) {

	val := ""
	hasPrefix := false

	if !strings.HasSuffix(flag, "=") {
		flag = flag + "="
	}

	for _, e := range args {
		hasPrefix = strings.HasPrefix(e, flag)
		if hasPrefix {
			val = e[len(flag):]
			break
		}
	}

	return val, hasPrefix
}

func (cf *CommandFlags) CurrentConfiguration() *EnvironmentConfig {
	if config == nil {
		config = cf.LoadConfig()
	}
	return config
}

func (env *EnvironmentConfig) String() string {
	js := enc.ToIndentedJson(env, "", "   ")
	return js
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
