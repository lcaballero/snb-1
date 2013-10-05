package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
)

func CreateTile(tileUuid, boardId, criteriaId string, position, active int) (codes.StatusCode, error) {

	var status codes.StatusCode

	sql := caching.Cache().CreateTile.Script

	_, err := sql_utils.GetConnection().Exec(
		string(sql),
		tileUuid, boardId, criteriaId, position, 1)

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else {
		status = codes.Success
		return status, err
	}

	return status, err

}
