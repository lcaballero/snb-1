package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
)

func CreateGame(gameUuid, groupId, name, description string) (codes.StatusCode, error) {

	var status codes.StatusCode

	sql := caching.CacheEntries.CreateGame.Script

	//gameUuid := uuid.New()

	// ?? should we assume groupId is valid or should
	// we run a db query to ensure it's valid?

	// winning_board_id == null
	// active is set to false to begin with until at least 25 criteria
	// are associated with the board
	_, err := sql_utils.GetConnection().Exec(
		sql, gameUuid, groupId, name, description, 1, 1)

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else {
		status = codes.Success
		return status, err
	}

	return status, err

}
