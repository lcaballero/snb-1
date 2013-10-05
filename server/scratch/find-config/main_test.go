package main

import (
	"testing"
)

func Test_Should_Find_Config_File(t *testing.T) {

	cf := "./config.js"

	_, ok := FindConfigFile(cf)

	if !ok {
		t.Error("Should have found the file: " + cf)
	}
}

func Test_Should_Not_Find_File(t *testing.T) {

	non := "./nononnnnnnnnn.js"

	_, ok := FindConfigFile(non)

	if ok {
		t.Error("Should not have found the file: " + non)
	}
}
