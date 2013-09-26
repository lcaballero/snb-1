package snap_sql

import (
	"fmt"
	"io/ioutil"
	"sql_utils"
	"sql_utils/codes"
)

func CreateTile(tileUuid, boardId, criteriaId string, position, active int) (codes.StatusCode, error) {

	var status codes.StatusCode

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createTile.sql")

	if err != nil {
		fmt.Println(err)
		status = codes.File_Read_Error
	} else {
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
	}

	return status, err

}
