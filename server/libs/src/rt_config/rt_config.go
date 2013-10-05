package rt_config

import (
	enc "json_helpers"
	"path"
	"path/filepath"
)

var (
	config *EnvironmentConfig = nil
)

// EnvironmentConfig holds the configuration variables for each of the
// environments as found in config.js
type EnvironmentConfig struct {
	ConfigFile string
	Dev        RuntimeConfig
	Acc        RuntimeConfig
	Prod       RuntimeConfig
}

// Contains parameters to pieces of the application, and is loaded via the
// EnvironmentConfig at runtime either by a command line flag or via a
// directory tree search for the file 'config.js' which lives in the root
// of the application directory.
type RuntimeConfig struct {
	ConnectionString string
	SqlScripts       string
	WebServerPort    string
	DbServerPort     string
}

// Access to the singleton of EnvironmentConfig which provides the
// RuntimConfig for the environment of choice.
func Config() *EnvironmentConfig {

	if config != nil {
		return config
	}

	config = newFlags().readCommandFlags().LoadEnvironmentConfig()

	return config
}

func findConfigFile(file string) string {

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

	return abs
}

func (r *RuntimeConfig) String() string {
	return enc.ToIndentedJson(r, "", "   ")
}

func (env *EnvironmentConfig) String() string {
	js := enc.ToIndentedJson(env, "", "   ")
	return js
}
