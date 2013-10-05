package models

import (
	"snap_sql"
	"sql_utils/caching"
	"testing"
)

func Test_Nothing(t *testing.T) {
	if caching.Cache() == nil {
		t.Error("The cache lookup was not found.")
	}
}

func Test_Initialization(t *testing.T) {
	snap_sql.DropAllTables()
	model := &InitialModel{}
	model.Initialize()
}
