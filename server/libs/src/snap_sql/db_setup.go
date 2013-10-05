package snap_sql

import (
	"fmt"
	"sql_utils"
)

const (
	SNB_DB = "snb"

	// These are the names of the tables as they appear in the database
	UserTable           = "_user"
	SocialGroupTable    = "socialgroup"
	UserToGroupTable    = "usertogroup"
	GameTable           = "game"
	BoardTable          = "board"
	CriteriaTable       = "criteria"
	TileTable           = "tile"
	GameToCriteriaTable = "gametocriteria"
)

var AllTables []string = []string{
	UserTable,
	SocialGroupTable,
	UserToGroupTable,
	GameTable,
	BoardTable,
	CriteriaTable,
	TileTable,
	GameToCriteriaTable,
}

func DropAllTables() {
	for _, table := range AllTables {
		ok := sql_utils.DropTable(table)
		if !ok {
			fmt.Println("Wasn't able to drop table: ", table)
		}
	}
}

func HasUserTable() bool {
	has_table := TableExists(UserTable)
	return has_table
}

func HasGroupTable() bool {
	has_table := TableExists(SocialGroupTable)
	return has_table
}

func HasUserToGroupTable() bool {
	has_table := TableExists(UserToGroupTable)
	return has_table
}

func HasGameTable() bool {
	return TableExists(GameTable)
}

func HasBoardTable() bool {
	return TableExists(BoardTable)
}

func HasCriteriaTable() bool {
	return TableExists(CriteriaTable)
}

func HasTileTable() bool {
	return TableExists(TileTable)
}

func HasGameToCriteriaTable() bool {
	return TableExists(GameToCriteriaTable)
}

func SetupTables() (err error) {

	if err == nil && !HasUserTable() {
		err = CreateUserTable()
	}

	if err == nil && !HasGroupTable() {
		err = CreateGroupsTable()
	}

	if err == nil && !HasUserToGroupTable() {
		err = CreateUserToGroupTable()
	}

	if err == nil && !HasGameTable() {
		err = CreateGameTable()
	}

	if err == nil && !HasBoardTable() {
		err = CreateBoardTable()
	}

	if err == nil && !HasCriteriaTable() {
		err = CreateCriteriaTable()
	}

	if err == nil && !HasTileTable() {
		err = CreateTileTable()
	}

	if err == nil && !HasGameToCriteriaTable() {
		err = CreateGameToCriteriaTable()
	}

	if err != nil {
		fmt.Println("Error setting up tables: ", err)
	}

	return nil
}
