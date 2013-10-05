package snap_sql

import (
	"fmt"
	"sql_utils"
	"testing"
	"uuid"
)

func Test_DoesFindRandomTable(t *testing.T) {

	table := uuid.New()

	has_table := sql_utils.TableExists(SNB_DB, table)

	if has_table {
		t.Error("Shouldn't have a " + table + " table")
	}
}

func Test_HasTable(t *testing.T) {

	dropped := sql_utils.DropTable("_user")
	has_table := HasUserTable()

	fmt.Println("dropped: ", dropped)

	if has_table {
		t.Error("Shouldn't have a _user table")
	}
}

func Test_CreateUserTable(t *testing.T) {

	has_table := HasUserTable()

	if !has_table {
		CreateUserTable()
	}

	has_table = sql_utils.TableExists("snb", "_user")

	if !has_table {
		t.Error("Should have created the _user table, but haven't")
	}
}

func Test_Dropped_All_Tables(t *testing.T) {

	DropAllTables()

	hasTables := HasUserTable() ||
		HasGroupTable() ||
		HasUserToGroupTable() ||
		HasGameTable() ||
		HasBoardTable() ||
		HasCriteriaTable() ||
		HasTileTable() ||
		HasGameToCriteriaTable()

	if hasTables {
		t.Error("All tables should be dropped, but are not.")
	}
}

func Test_Setup_Db(t *testing.T) {

	SetupTables()

	if !HasUserTable() {
		t.Error("Failed to create user table.")
	}

	if !HasGroupTable() {
		t.Error("Failed to create group table.")
	}

	if !HasUserToGroupTable() {
		t.Error("Failed to create a User to Group table")
	}

	if !HasGameTable() {
		t.Error("Failed to create a Game table.")
	}

	if !HasBoardTable() {
		t.Error("Failed to create the Board Table.")
	}

	if !HasCriteriaTable() {
		t.Error("Failed to create the Criteria Table.")
	}

	if !HasTileTable() {
		t.Error("Failed to create the Tile Table.")
	}

	if !HasGameToCriteriaTable() {
		t.Error("Failed to create the GameToCriteria Table.")
	}
}
