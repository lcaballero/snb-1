package snap_sql

import(
	"fmt"
	"io/ioutil"
	"sql_utils"
)

func CreateGroup(groupUuid, group_name, group_desc, group_owner string) (sql_utils.StatusCode, error) {
		
	has_group, err := HasGroup(group_name)

	var status sql_utils.StatusCode;

	if err != nil {
		fmt.Println(err)
		status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
	} else if has_group {
		status = sql_utils.STATUS_CODES[sql_utils.GROUP_EXISTS]
	} else {
		sql, err := ioutil.ReadFile(sql_utils.FilePath + "createGroup.sql")

		if err != nil {
			fmt.Println(err)
			status = sql_utils.STATUS_CODES[sql_utils.FILE_READ_ERR]
		} else {
			
			_, err := sql_utils.GetConnection().Exec(string(sql), groupUuid, group_name, group_desc, group_owner)

			if err != nil {
				fmt.Println(err)
				status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
			} else {
				status = sql_utils.STATUS_CODES[sql_utils.SUCCESS]
				return status, err
			}
		}
	}

	return status, err
}