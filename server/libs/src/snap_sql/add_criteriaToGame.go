package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
	"uuid"
)

func AddCriteriaToGame(game_id, criteria_id string) (codes.StatusCode, error) {

	var status codes.StatusCode

	sql := caching.CacheEntries.AddCriteriaToGame.Script

	rowUuid := uuid.New()
	_, err := sql_utils.GetConnection().Exec(sql, rowUuid, game_id, criteria_id, 1, 1)

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else {
		status = codes.Success
		return status, err
	}

	fmt.Println(err)
	return status, err
}
