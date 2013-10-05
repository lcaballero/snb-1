package models

import (
	"fmt"
	"sql_utils/caching"
	"testing"
)

func Test_Nothing(t *testing.T) {
	fmt.Println("here")

	if caching.Cache() == nil {
		t.Error("The cache lookup was not found.")
	}
}

func Test_Initialization(t *testing.T) {
	model := &InitialModel{}
	model.Initialize()
}
