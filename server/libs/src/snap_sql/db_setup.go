package snap_sql

import (
	"fmt"
	"sql_utils"
)

const (
	SNB_DB              = "snb"
	UserTable           = "_user"
	SocialGroupTable    = "socialgroup"
	UserToGroupTable    = "usertogroup"
	GameTable           = "game"
	BoardTable          = "board"
	CriteriaTable       = "criteria"
	TileTable           = "tile"
	GameToCriteriaTable = "gametocriteria"
)

func DropAllTables() {
	sql_utils.DropTable(UserTable)
	sql_utils.DropTable(SocialGroupTable)
	sql_utils.DropTable(UserToGroupTable)
	sql_utils.DropTable(GameTable)
	sql_utils.DropTable(BoardTable)
	sql_utils.DropTable(CriteriaTable)
	sql_utils.DropTable(TileTable)
	sql_utils.DropTable(GameToCriteriaTable)
}

func HasUserTable() bool {
	has_table := sql_utils.TableExists(SNB_DB, UserTable)
	return has_table
}

func HasGroupTable() bool {
	has_table := sql_utils.TableExists(SNB_DB, SocialGroupTable)
	return has_table
}

func HasUserToGroupTable() bool {
	has_table := sql_utils.TableExists(SNB_DB, UserToGroupTable)
	return has_table
}

func HasGameTable() bool {
	return sql_utils.TableExists(SNB_DB, GameTable)
}

func HasBoardTable() bool {
	return sql_utils.TableExists(SNB_DB, BoardTable)
}

func HasCriteriaTable() bool {
	return sql_utils.TableExists(SNB_DB, CriteriaTable)
}

func HasTileTable() bool {
	return sql_utils.TableExists(SNB_DB, TileTable)
}

func HasGameToCriteriaTable() bool {
	return sql_utils.TableExists(SNB_DB, GameToCriteriaTable)
}

func SetupTables() (err error) {

	if err == nil && !HasUserTable() {
		fmt.Println("Creating User Table...")
		err = CreateUserTable()
	}

	if err == nil && !HasGroupTable() {
		fmt.Println("Creating Groups Table...")
		err = CreateGroupsTable()
	}

	if err == nil && !HasUserToGroupTable() {
		fmt.Println("Creating User To Group Table...")
		err = CreateUserToGroupTable()
	}

	if err == nil && !HasGameTable() {
		fmt.Println("Creating Game Table...")
		err = CreateGameTable()
	}

	if err == nil && !HasBoardTable() {
		fmt.Println("Creating Board Table...")
		err = CreateBoardTable()
	}

	if err == nil && !HasCriteriaTable() {
		fmt.Println("Creating Criteria Table...")
		err = CreateCriteriaTable()
	}

	if err == nil && !HasTileTable() {
		fmt.Println("Creating Tile Table...")
		err = CreateTileTable()
	}

	if err == nil && !HasGameToCriteriaTable() {
		fmt.Println("Creating GameToCriteria Table...")
		err = CreateGameToCriteriaTable()
	}

	if err != nil {
		fmt.Println("Error setting up tables: ", err)
	}

	return nil
}
