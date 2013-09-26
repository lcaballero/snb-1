package rt_config

import (
	"path"
	"testing"
)

func Test_Initial_Config_State(t *testing.T) {

	if config != nil {
		t.Error("Private 'config' should start as nil.")
	}
}

func Test_Default_Config_File(t *testing.T) {

	file := defaultFileProvider()

	if file == "./config.js" {
		default_file := path.Join(default_config_root, default_config_name)
		t.Error("Default path should be: ", default_file)
	}
}
