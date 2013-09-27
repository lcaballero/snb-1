package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
)

func CreateCriteria(criteriaUuid, description string) (codes.StatusCode, error) {

	var status codes.StatusCode

	sql := caching.CacheEntries.CreateCriteria.Script

	//gameUuid := uuid.New()

	// ?? should we assume groupId is valid or should
	// we run a db query to ensure it's valid?

	_, err := sql_utils.GetConnection().Exec(sql, criteriaUuid, description, 1)

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else {
		status = codes.Success
		return status, err
	}

	return status, err

}
