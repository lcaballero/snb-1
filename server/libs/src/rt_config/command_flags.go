package rt_config

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
