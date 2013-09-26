package snap_sql

import (
	"fmt"
	"io/ioutil"
	"sql_utils"
)

func CreateTile(tileUuid, boardId, criteriaId string, position, active int) (sql_utils.StatusCode, error) {

	var status sql_utils.StatusCode

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createTile.sql")

	if err != nil {
		fmt.Println(err)
		status = sql_utils.STATUS_CODES[sql_utils.FILE_READ_ERR]
	} else {
		_, err := sql_utils.GetConnection().Exec(
			string(sql),
			tileUuid, boardId, criteriaId, position, 1)

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
