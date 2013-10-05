package rt_config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	enc "json_helpers"
	"os"
	"path"
	"strings"
)

type EnvironmentConfig struct {
	ConfigFile     string
	Dev, Acc, Prod RuntimeConfig
}

type RuntimeConfig struct {
	ConnectionString string
	SqlScripts       string
	WebServerPort    string
	DbServerPort     string
}

type ConfigFileProvider func() string

type CommandFlags struct {
	ConfigFile string
	SqlScripts string
}

var config *EnvironmentConfig = nil
var PathProvider ConfigFileProvider = defaultFileProvider

const (
	default_config_name = "config.js"
	default_config_root = "./"
	flag_config_file    = "--config-file"
	flag_sql_scripts    = "--sql-scripts"
)

func LoadFromCommandLine() *EnvironmentConfig {

	cf := readFlags()

	pwd, e := os.Getwd()

	fmt.Println("pwd: ", pwd, e)
	fmt.Println("cf.ConfigFile ", cf.ConfigFile)

	_, err := os.Stat(cf.ConfigFile)

	configFileExists := cf.ConfigFile != "" && !os.IsNotExist(err)

	if configFileExists {

		PathProvider = func() string { return cf.ConfigFile }
		return CurrentConfiguration()

	} else {

		fmt.Println("Configuration doesn't exist: ", cf.ConfigFile)
		fmt.Println(enc.ToIndentedJson(cf, " ", "   "))
	}

	return &EnvironmentConfig{}
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

func FindConfigFile(file string) (string, bool) {

	exists := func(file string) bool {
		_, err := os.Stat(file)
		exists := file != "" && !os.IsNotExist(err)
		return exists
	}

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

func readFlags() *CommandFlags {

	cf := &CommandFlags{}

	cf.ConfigFile, _ = FindFlag(flag_config_file, os.Args)
	cf.SqlScripts, _ = FindFlag(flag_sql_scripts, os.Args)

	return cf
}

func CurrentConfiguration() *EnvironmentConfig {
	if config == nil {
		config = LoadConfig(PathProvider())
	}
	return config
}

func defaultFileProvider() string {
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
	val.ConfigFile = path

	if err != nil {
		fmt.Println(err)
	}

	return val
}
