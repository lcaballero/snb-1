package snap_sql

import (
	"data_classes"
	"database/sql"
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	"time"
	//enc "json_helpers"
)

// ---------------------- Read Group Functions ---------------------- //

func HasBoard(id string) (bool, error) {
	board, err := ReadBoard(id)

	if err != nil {
		fmt.Println(err)
		return true, err // TODO: should this be true or false?
	} else if len(board) > 0 {
		return true, err
	} else {
		return false, err
	}
}

func ReadBoard(id string) ([]data_classes.BoardData, error) {
	//sql := "SELECT * FROM board WHERE id=$1"
	sql := caching.CacheEntries.ReadBoardFromId.Script

	return processBoard(sql_utils.GetConnection().Query(sql, id))
}

func ReadUsersBoards(userId string) ([]data_classes.BoardData, error) {
	// sql := SELECT * FROM board WHERE user_id = $1
	sql := caching.CacheEntries.ReadUsersBoards.Script
	return processBoard(sql_utils.GetConnection().Query(sql, userId))
}

func processBoard(sqlRows *sql.Rows, err error) ([]data_classes.BoardData, error) {
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		mappedRows := sql_utils.ToSqlMap(sqlRows)

		boards := make([]data_classes.BoardData, len(mappedRows))

		for i, v := range mappedRows {
			u := data_classes.BoardData{
				Id:        v["id"].(string),
				UserId:    v["user_id"].(string),
				GameId:    v["game_id"].(string),
				Name:      v["name"].(string),
				State:     v["state"].(int64),
				Active:    v["active"].(int64),
				DateAdded: v["date_added"].(time.Time),
			}

			boards[i] = u
		}

		return boards, nil
	}
}
