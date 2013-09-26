package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
)

func CreateBoard(boardUuid, gameId, userId, name string, state int) (codes.StatusCode, error) {

	var status codes.StatusCode

	sql := caching.CacheEntries.CreateBoard.Script

	//gameUuid := uuid.New()

	// ?? should we assume groupId is valid or should
	// we run a db query to ensure it's valid?

	_, err := sql_utils.GetConnection().Exec(
		string(sql),
		boardUuid, gameId, userId, name, state)

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else {
		status = codes.Success
		return status, err
	}

	return status, err

}
