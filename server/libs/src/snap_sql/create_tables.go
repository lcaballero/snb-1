package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
)

func TableExists(name string) bool {
	return sql_utils.TableExists(SNB_DB, name)
}

func CreateGameToCriteriaTable() (err error) {

	sql := caching.Cache().CreateGameToCriteriaTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println("Creating Tile Table Error:", err)
	}

	return err
}

func CreateTileTable() (err error) {

	sql := caching.Cache().CreateTileTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println("Creating Tile Table Error:", err)
	}

	return err
}

func CreateCriteriaTable() (err error) {

	sql := caching.Cache().CreateCriteriaTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println("Creating Criteria Table:", err)
	}

	return err
}

func CreateBoardTable() (err error) {

	sql := caching.Cache().CreateBoardTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func CreateGameTable() (err error) {

	sql := caching.Cache().CreateGameTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func CreateUserTable() (err error) {

	sql := caching.Cache().CreateUserTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

func CreateGroupsTable() (err error) {

	sql := caching.Cache().CreateGroupTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println("CreateGroupsTable", err)
	}

	return err
}

func CreateUserToGroupTable() (err error) {

	sql := caching.Cache().CreateUserToGroupTable.Script

	_, err = sql_utils.GetConnection().Exec(sql)

	if err != nil {
		fmt.Println(err)
	}

	return err
}
