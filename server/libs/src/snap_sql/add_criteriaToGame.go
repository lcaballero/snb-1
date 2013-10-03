package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
	//"uuid"
)

func AddCriteriaToGame(uuid, game_id, criteria_id string, state, active int) (codes.StatusCode, error) {

	var status codes.StatusCode

	sql := caching.CacheEntries.AddCriteriaToGame.Script

	//rowUuid := uuid.New()
	_, err := sql_utils.GetConnection().Exec(sql, uuid, game_id, criteria_id, state, active)

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else {
		status = codes.Success
		return status, err
	}

	return status, err
}
