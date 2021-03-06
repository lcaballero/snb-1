package snap_sql

import (
	logging "log"
	"os"
	"sql_utils"
	"testing"
	"uuid"
)

var log *logging.Logger = nil

func init() {
	log = logging.New(os.Stdout, "", logging.Lshortfile)

	log.Println("Init-ing the log.")
}

func Test_DoesFindRandomTable(t *testing.T) {

	log.Println("Test_DoesFindRandomTable.")

	table := uuid.New()

	has_table := sql_utils.TableExists(SNB_DB, table)

	if has_table {
		t.Error("Shouldn't have a " + table + " table")
	}
}

// func Test_HasTable(t *testing.T) {

// 	sql_utils.DropTable(UserTable)

// 	has_table := HasUserTable()

// 	if has_table {
// 		t.Error("Shouldn't have a _user table")
// 	}
//}

func Test_CreateUserTable(t *testing.T) {

	has_table := HasUserTable()

	if !has_table {
		CreateUserTable()
	}

	has_table = sql_utils.TableExists(SNB_DB, UserTable)

	if !has_table {
		t.Error("Should have created the _user table, but hasn't")
	}
}

func Test_Dropped_All_Tables(t *testing.T) {

	DropAllTables()

	created := map[string]bool{
		UserTable:           HasUserTable(),
		SocialGroupTable:    HasGroupTable(),
		UserToGroupTable:    HasUserToGroupTable(),
		GameTable:           HasGameTable(),
		BoardTable:          HasBoardTable(),
		CriteriaTable:       HasCriteriaTable(),
		TileTable:           HasTileTable(),
		GameToCriteriaTable: HasGameToCriteriaTable(),
	}

	msg := ""

	for table, hasTable := range created {
		if hasTable {
			msg = msg + "\nDidn't drop table: " + table
		}
	}

	if msg != "" {
		t.Error(msg)
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
