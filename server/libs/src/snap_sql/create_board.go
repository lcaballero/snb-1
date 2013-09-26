package snap_sql

import (
	"fmt"
	"io/ioutil"
	"sql_utils"
	"sql_utils/codes"
)

func CreateBoard(boardUuid, gameId, userId, name string, state int) (codes.StatusCode, error) {

	var status codes.StatusCode

	sql, err := ioutil.ReadFile(sql_utils.FilePath + "createBoard.sql")

	if err != nil {
		fmt.Println(err)
		status = codes.File_Read_Error
	} else {
		//gameUuid := uuid.New()

		// ?? should we assume groupId is valid or should
		// we run a db query to ensure it's valid?

		_, err := sql_utils.GetConnection().Exec(
			string(sql),
			boardUuid, gameId, userId, name, state)

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
