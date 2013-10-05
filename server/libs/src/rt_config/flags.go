package rt_config

import (
	"os"
	"strings"
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

func (f *flags) readCommandFlags() *CommandFlags {

	cf := &CommandFlags{}

	cf.ConfigFile, _ = FindFlag(f.configFile, os.Args)
	cf.SqlScripts, _ = FindFlag(f.sqlScripts, os.Args)

	return cf
}

func newFlags() *flags {
	return &flags{
		configFile: flag_config_file,
		sqlScripts: flag_sql_scripts,
	}
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
