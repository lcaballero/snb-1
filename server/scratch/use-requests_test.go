package main

import (
	"fmt"
	"snap_sql"
	"testing"
	"uuid"
)

func Test_DoesFindRandomTable(t *testing.T) {
	table := uuid.New()
	has_table := tableExists("snb", table)

	if has_table {
		t.Error("Shouldn't have a " + table + " table")
	}
}

func Test_HasTable(t *testing.T) {

	dropped := dropTable("_user")
	has_table := tableExists("snb", "_user")

	fmt.Println("dropped: ", dropped)

	if has_table {
		t.Error("Shouldn't have a _user table")
	}
}

func Test_CreateUserTable(t *testing.T) {
	has_table := tableExists("snb", "_user")

	if !has_table {
		snap_sql.CreateUserTable()
	}

	has_table = tableExists("snb", "_user")

	if !has_table {
		t.Error("Should have created the _user table, but haven't")
	}
}
