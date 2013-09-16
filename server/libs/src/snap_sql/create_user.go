package snap_sql

import(
	"fmt"
	"uuid"
	"io/ioutil"
	"sql_utils"
)

func CreateUser(email, password string) (sql_utils.StatusCode, error) {

	has_user, err := HasUser(email)

	var status sql_utils.StatusCode;

	if err != nil {
		fmt.Println(err)
		status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
	} else if has_user {
		status = sql_utils.STATUS_CODES[sql_utils.USER_EXISTS]
	} else {

		sql, err := ioutil.ReadFile(sql_utils.FilePath+"createUser.sql")

		if err != nil {
			fmt.Println(1, err)
			status = sql_utils.STATUS_CODES[sql_utils.FILE_READ_ERR]
		} else {
			userUuid := uuid.New()
			_, err := sql_utils.GetConnection().Exec(string(sql), userUuid, email, password)
			
			if err != nil {
				fmt.Println(2, err)
				status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
			} else {
				//fmt.Println("Create User result: ", result)

				group, err := ReadGroup("global_group")

				if err != nil {
					status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
				} else {

					status, err := AddUserToGroup(userUuid, group[0].Id)
					fmt.Println(status.Msg)

					if err != nil{
						fmt.Println(err)
						
					} else {
						status = sql_utils.STATUS_CODES[sql_utils.SUCCESS]
						return status, err
					}
				}
			}
		
		}
	}

	return status, err
}