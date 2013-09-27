package snap_sql

import (
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"sql_utils/codes"
	"uuid"
)

func AddUserToGroup(userId, groupId string) (codes.StatusCode, error) {

	var status codes.StatusCode
	// has_group, err := hasGroup(groupId)

	// if has_group {

	// has_user, err := hasUserId(userId)

	//if has_user && err == nil {
	sql := caching.CacheEntries.AddUserToGroup.Script

	// add user to the global group
	rowUuid := uuid.New()
	_, err := sql_utils.GetConnection().Exec(sql, rowUuid, groupId, userId)

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else {
		status = codes.Success
		return status, err
	}

	// } else {
	// 	status = STATUS_CODES[USER_DOES_NOT_EXISTS]
	// }
	// } else {
	// 	status = STATUS_CODES[GROUP_DOES_NOT_EXISTS]
	// }

	fmt.Println(err)
	return status, err
}
