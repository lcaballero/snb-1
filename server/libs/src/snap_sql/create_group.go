package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
)

func CreateGroup(groupUuid, group_name, group_desc, group_owner string) (codes.StatusCode, error) {

	has_group, err := HasGroup(group_name)

	var status codes.StatusCode

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else if has_group {
		status = codes.Group_Exists
	} else {
		sql := caching.CacheEntries.CreateGroup.Script

		_, err := sql_utils.GetConnection().Exec(sql, groupUuid, group_name, group_desc, group_owner)

		if err != nil {
			fmt.Println(err)
			status = codes.Db_Error
		} else {
			status = codes.Success
			return status, err
		}
	}

	return status, err
}
