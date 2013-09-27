package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
)

// ---------------------- Create Tile Table ---------------------- //

func CreateTileTable() {

	sql := caching.CacheEntries.CreateTileTable.Script

	_, err := sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}
}

// ---------------------- Create Criteria Table ---------------------- //

func CreateCriteriaTable() {

	sql := caching.CacheEntries.CreateCriteriaTable.Script

	_, err := sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}
}

// ---------------------- Create Board Table ---------------------- //

func CreateBoardTable() (err error) {

	sql := caching.CacheEntries.CreateBoardTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

// ---------------------- Create Game Table ---------------------- //

func CreateGameTable() (err error) {

	sql := caching.CacheEntries.CreateGameTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

// ---------------------- Create User Table ---------------------- //
func CreateUserTable() (err error) {

	sql := caching.CacheEntries.CreateUserTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

// ---------------------- Create Group Table ---------------------- //

func CreateGroupsTable() (err error) {

	sql := caching.CacheEntries.CreateGroupTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

// ---------------------- Create User to Group table ---------------------- //

func CreateUserToGroupTable() (err error) {

	sql := caching.CacheEntries.CreateUserToGroupTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
