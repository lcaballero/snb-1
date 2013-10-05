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

func ReadTile(id string) ([]data_classes.TileData, error) {

	sql := caching.Cache().ReadTile.Script

	return processTile(sql_utils.GetConnection().Query(sql, id))
}

func ReadBoardTiles(boardId string) ([]data_classes.TileData, error) {

	sql := caching.Cache().ReadBoardTiles.Script

	return processTile(sql_utils.GetConnection().Query(sql, boardId))
}

func processTile(sqlRows *sql.Rows, err error) ([]data_classes.TileData, error) {
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		mappedRows := sql_utils.ToSqlMap(sqlRows)

		data := make([]data_classes.TileData, len(mappedRows))

		for i, v := range mappedRows {
			u := data_classes.TileData{
				Id:         v["id"].(string),
				BoardId:    v["board_id"].(string),
				CriteriaId: v["criteria_id"].(string),
				Position:   v["position"].(int64),
				State:      v["state"].(int64),
				Active:     v["active"].(int64),
				DateAdded:  v["date_added"].(time.Time),
			}

			data[i] = u
		}

		return data, nil
	}
}
