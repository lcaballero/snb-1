package snap_sql

import (
	"fmt"
	"io/ioutil"
	"sql_utils"
	"sql_utils/codes"
	"uuid"
)

func CreateUser(email, password string) (codes.StatusCode, error) {

	has_user, err := HasUser(email)

	var status codes.StatusCode

	if err != nil {
		fmt.Println(err)
		status = codes.Db_Error
	} else if has_user {
		status = codes.User_Exists
	} else {

		sql, err := ioutil.ReadFile(sql_utils.FilePath + "createUser.sql")

		if err != nil {
			fmt.Println(err)
			status = codes.File_Read_Error
		} else {
			userUuid := uuid.New()
			_, err := sql_utils.GetConnection().Exec(string(sql), userUuid, email, password)

			if err != nil {
				fmt.Println(err)
				status = codes.Db_Error
			} else {
				//fmt.Println("Create User result: ", result)

				group, err := ReadGroup("global_group")

				if err != nil {
					status = codes.Db_Error
				} else {

					status, err := AddUserToGroup(userUuid, group[0].Id)
					fmt.Println(status.Msg)

					if err != nil {
						fmt.Println(err)

					} else {
						status = codes.Success
						return status, err
					}
				}
			}

		}
	}

	return status, err
}
