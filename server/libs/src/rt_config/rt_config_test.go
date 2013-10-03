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

var fake_args []string = []string{
	"--config-file=../../",
	"--sql-scripts=../sqlScripts/",
}

func Test_Finds_Flag_Without_Equals(t *testing.T) {

	val, ok := FindFlag("--config-file", fake_args)

	if !ok {
		t.Error("Should have found the config-file flag.")
	}

	if val != "../../" {
		t.Error("Didn't find right value for the flag: --config-file, instead found: " + val)
	}
}

func Test_Finds_Flag_With_Equals(t *testing.T) {

	val, ok := FindFlag("--config-file=", fake_args)

	if !ok {
		t.Error("Should have found the config-file flag.")
	}

	if val != "../../" {
		t.Error("Didn't find right value for the flag: --config-file, instead found: " + val)
	}
}

func Test_Handles_Nil_Args(t *testing.T) {

	_, ok := FindFlag("--config-file=", nil)

	if ok {
		t.Error("Should not have found any values.")
	}
}
