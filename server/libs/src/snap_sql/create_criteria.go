package snap_sql

import (
	"fmt"
	"io/ioutil"
	"sql_utils"
)

func CreateCriteria(criteriaUuid, description string) (sql_utils.StatusCode, error) {

	var status sql_utils.StatusCode

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createCriteria.sql")

	if err != nil {
		fmt.Println(err)
		status = sql_utils.STATUS_CODES[sql_utils.FILE_READ_ERR]
	} else {
		//gameUuid := uuid.New()

		// ?? should we assume groupId is valid or should
		// we run a db query to ensure it's valid?

		_, err := sql_utils.GetConnection().Exec(
			string(sql),
			criteriaUuid, description, 1)

		if err != nil {
			fmt.Println(err)
			status = sql_utils.STATUS_CODES[sql_utils.DB_ERR]
		} else {
			status = sql_utils.STATUS_CODES[sql_utils.SUCCESS]
			return status, err
		}
	}

	return status, err

}
