package snap_sql

import (
	"fmt"
	"sql_utils"
)

func HasUserTable() bool {
	has_table := sql_utils.TableExists("snb", "_user")
	return has_table
}

func HasGroupTable() bool {
	has_table := sql_utils.TableExists("snb", "socialgroup")
	return has_table
}

func HasUserToGroupTable() bool {
	has_table := sql_utils.TableExists("snb", "usertogroup")
	return has_table
}

func HasGameTable() bool {
	return sql_utils.TableExists("snb", "game")
}

func HasBoardTable() bool {
	return sql_utils.TableExists("snb", "board")
}

func HasCriteriaTable() bool {
	return sql_utils.TableExists("snb", "criteria")
}

func HasTileTable() bool {
	return sql_utils.TableExists("snb", "tile")
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

	fmt.Println("Setting up Tile.")

	if err == nil && !HasTileTable() {
		fmt.Println("Creating Tile Table...")
		err = CreateTileTable()
	}

	if err != nil {
		fmt.Println("Error setting up tables: ", err)
	}

	return nil
}
