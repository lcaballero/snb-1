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

func SetupTables() (err error) {

	if err != nil && !HasUserTable() {
		err = CreateUserTable()
	}

	if err != nil && !HasGroupTable() {
		err = CreateGroupsTable()
	}

	if err != nil && !HasUserToGroupTable() {
		err = CreateUserToGroupTable()
	}

	if err != nil && !HasGameTable() {
		err = CreateGameTable()
	}

	if err != nil && !HasBoardTable() {
		err = CreateBoardTable()
	}

	if err == nil {
		fmt.Println(err)
	}

	return nil
}
