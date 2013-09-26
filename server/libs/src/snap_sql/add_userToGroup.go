package snap_sql

import (
	"fmt"
	"io/ioutil"
	"sql_utils"
	"uuid"
)

func AddUserToGroup(userId, groupId string) (sql_utils.StatusCode, error) {

	var status sql_utils.StatusCode
	// has_group, err := hasGroup(groupId)

	// if has_group {

	// has_user, err := hasUserId(userId)

	//if has_user && err == nil {
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "addUserToGroup.sql")

	if err != nil {
		fmt.Println(err)
		status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
	} else {

		// add user to the global group
		rowUuid := uuid.New()
		_, err := sql_utils.GetConnection().Exec(string(sql), rowUuid, groupId, userId)

		if err != nil {
			fmt.Println(err)
			status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
		} else {
			status = sql_utils.STATUS_CODES[sql_utils.SUCCESS]
			return status, err
		}
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
