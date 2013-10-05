package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
)

// ---------------------- Create Tile Table ---------------------- //

func CreateGameToCriteriaTable() (err error) {

	sql := caching.Cache().CreateGameToCriteriaTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println("Creating Tile Table Error:", err)
	}

	return err
}

// ---------------------- Create Tile Table ---------------------- //

func CreateTileTable() (err error) {

	sql := caching.Cache().CreateTileTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println("Creating Tile Table Error:", err)
	}

	return err
}

// ---------------------- Create Criteria Table ---------------------- //

func CreateCriteriaTable() (err error) {

	sql := caching.Cache().CreateCriteriaTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println("Creating Criteria Table:", err)
	}

	return err
}

// ---------------------- Create Board Table ---------------------- //

func CreateBoardTable() (err error) {

	sql := caching.Cache().CreateBoardTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

// ---------------------- Create Game Table ---------------------- //

func CreateGameTable() (err error) {

	sql := caching.Cache().CreateGameTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

// ---------------------- Create User Table ---------------------- //
func CreateUserTable() (err error) {

	sql := caching.Cache().CreateUserTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

// ---------------------- Create Group Table ---------------------- //

func CreateGroupsTable() (err error) {

	sql := caching.Cache().CreateGroupTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println("CreateGroupsTable", err)
	}

	return err
}

// ---------------------- Create User to Group table ---------------------- //

func CreateUserToGroupTable() (err error) {

	sql := caching.Cache().CreateUserToGroupTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}
