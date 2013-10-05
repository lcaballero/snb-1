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

func ReadCriteria(id string) ([]data_classes.CriteriaData, error) {
	//sql := "SELECT * FROM board WHERE id=$1"
	sql := caching.Cache().ReadCriteriaFromId.Script

	return processCriteria(sql_utils.GetConnection().Query(sql, id))
}

/*
func ReadBoardCriteria(boardId string) ([]data_classes.BoardData, error) {
	// sql := SELECT * FROM board WHERE user_id = $1
	sql := caching.CacheEntries.ReadBoardCriteria.Script
	return processBoard(sql_utils.GetConnection().Query(sql, userId))
}
*/
func processCriteria(sqlRows *sql.Rows, err error) ([]data_classes.CriteriaData, error) {
	if err != nil {
		fmt.Println("processCriteria", err)
		return nil, err
	} else {

		mappedRows := sql_utils.ToSqlMap(sqlRows)

		crit := make([]data_classes.CriteriaData, len(mappedRows))

		for i, v := range mappedRows {
			u := data_classes.CriteriaData{
				Id:        v["id"].(string),
				UserId:    v["user_id"].(string),
				GameId:    v["game_id"].(string),
				Name:      v["name"].(string),
				DateAdded: v["date_added"].(time.Time),
			}

			crit[i] = u
		}

		return crit, nil
	}
}
