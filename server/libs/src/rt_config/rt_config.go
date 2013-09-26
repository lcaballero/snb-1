package rt_config

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	enc "json_helpers"
	"os"
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

type ConfigFileProvider func() string

type CommandFlags struct {
	config_file     string
	help            bool
	show_parameters bool
}

var config *EnvironmentConfig = nil

const (
	default_config_name = "config.js"
	default_config_root = "./"
)

var PathProvider ConfigFileProvider = defaultFileProvider

func (r *RuntimeConfig) String() string {
	return enc.ToIndentedJson(r, "", "   ")
}

func LoadFromCommandLine() *EnvironmentConfig {

	cf := readFlags()

	fmt.Println("config_file: ", cf.config_file)

	if cf.help {
		flag.Usage()
		return &EnvironmentConfig{}
	}

	_, err := os.Stat(cf.config_file)

	configFileExists := !os.IsNotExist(err)

	if cf.config_file != "" && configFileExists {
		PathProvider = func() string {
			return cf.config_file
		}
		return CurrentConfiguration()
	} else {
		fmt.Println("Configuration doesn't exist: ", cf.config_file)
	}

	return &EnvironmentConfig{}
}

func readFlags() *CommandFlags {

	cf := &CommandFlags{}

	flag.StringVar(&cf.config_file, "config-file", "", "Absolute path to the root of the application where the config.js should reside.")
	flag.BoolVar(&cf.help, "help", false, "Show this help message.")

	flag.Parse()

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

	if err != nil {
		fmt.Println(err)
	}

	return val
}
