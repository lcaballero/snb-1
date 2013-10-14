package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
)

func UpdateBoardState(id string, state int) (codes.StatusCode, error) {
	return updateState(id, "board", state)
}

func UpdateTileState(id string, state int) (codes.StatusCode, error) {
	return updateState(id, "tile", state)
}

func UpdateGameState(id string, state int) (codes.StatusCode, error) {
	return updateState(id, "game", state)
}

func UpdateGroupState(id string, state int) (codes.StatusCode, error) {
	return updateState(id, "group", state)
}

func UpdateCriteriaState(id string, state int) (codes.StatusCode, error) {
	return updateState(id, "criteria", state)
}

func updateState(id, table string, state int) (codes.StatusCode, error) {

	var status codes.StatusCode

	var sql string
	if table == "board" {
		sql = caching.Cache().UpdateBoardState.Script
	} else if table == "tile" {
		sql = caching.Cache().UpdateTileState.Script
	} else if table == "game" {
		sql = caching.Cache().UpdateGameState.Script
	} else if table == "group" {
		sql = caching.Cache().UpdateGroupState.Script
	} else if table == "criteria" {
		sql = caching.Cache().UpdateCriteriaState.Script
	}

	fmt.Println("sql...........", sql)

	_, err := sql_utils.GetConnection().Exec(
		sql, state, id)

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else {
		status = codes.Success
		return status, err
	}

	return status, err
}

func ActivateBoard(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "board", 1)
}

func ActivateTile(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "tile", 1)
}

func ActivateGame(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "game", 1)
}

func ActivateGroup(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "group", 1)
}

func ActivateCriteria(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "criteria", 1)
}

func deactivateBoard(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "board", 0)
}

func deactivateTile(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "tile", 0)
}

func deactivateGame(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "game", 0)
}

func deactivateGroup(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "group", 0)
}

func deactivateCriteria(id string) (codes.StatusCode, error) {
	return updateActiveFlag(id, "criteria", 0)
}

func updateActiveFlag(id, table string, state int) (codes.StatusCode, error) {

	var status codes.StatusCode

	var sql string
	if table == "board" {
		sql = caching.Cache().UpdateBoardActiveFlag.Script
	} else if table == "tile" {
		sql = caching.Cache().UpdateTileActiveFlag.Script
	} else if table == "game" {
		sql = caching.Cache().UpdateGameActiveFlag.Script
	} else if table == "group" {
		sql = caching.Cache().UpdateGroupActiveFlag.Script
	} else if table == "criteria" {
		sql = caching.Cache().UpdateCriteriaActiveFlag.Script
	}

	fmt.Println("sql...........", sql)

	_, err := sql_utils.GetConnection().Exec(
		sql, state, id)

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else {
		status = codes.Success
		return status, err
	}

	return status, err
}
