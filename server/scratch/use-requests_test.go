package main

import (
	"fmt"
	"snap_sql"
	"sql_utils"
	"testing"
	"uuid"
)

func Test_DoesFindRandomTable(t *testing.T) {
	table := uuid.New()
	has_table := sql_utils.TableExists("snb", table)

	if has_table {
		t.Error("Shouldn't have a " + table + " table")
	}
}

func Test_HasTable(t *testing.T) {

	dropped := sql_utils.DropTable("_user")
	has_table := sql_utils.TableExists("snb", "_user")

	fmt.Println("dropped: ", dropped)

	if has_table {
		t.Error("Shouldn't have a _user table")
	}
}

func Test_CreateUserTable(t *testing.T) {
	has_table := sql_utils.TableExists("snb", "_user")

	if !has_table {
		snap_sql.CreateUserTable()
	}

	has_table = sql_utils.TableExists("snb", "_user")

	if !has_table {
		t.Error("Should have created the _user table, but haven't")
	}
}

func Test_Setup_Db(t *testing.T) {
	snap_sql.SetupTables()

	if !snap_sql.HasUserTable() {
		t.Error("Failed to create user table.")
	}

	if !snap_sql.HasGroupTable() {
		t.Error("Failed to create group table.")
	}

	if !snap_sql.HasUserToGroupTable() {
		t.Error("Failed to create a User to Group table")
	}

	if !snap_sql.HasGameTable() {
		t.Error("Failed to create a Game table.")
	}

	if !snap_sql.HasBoardTable() {
		t.Error("Failed to create the Board Table.")
	}
}
