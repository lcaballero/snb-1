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

func ReadGameFromId(gameId string) ([]data_classes.GameData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql := caching.Cache().ReadGameFromId.Script

	return processGames(sql_utils.GetConnection().Query(sql, gameId))
}

func ReadGameFromName(gameName string) ([]data_classes.GameData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"

	sql := caching.Cache().ReadGameFromName.Script

	return processGames(sql_utils.GetConnection().Query(sql, gameName))
}

func ReadGameInGroupFromName(groupId, gameName string) ([]data_classes.GameData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql := caching.Cache().ReadGameInGroupFromName.Script

	return processGames(sql_utils.GetConnection().Query(sql, groupId, gameName))
}

func ReadAllGames(groupId string) ([]data_classes.GameData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql := caching.Cache().ReadAllGamesInGroup.Script

	return processGames(sql_utils.GetConnection().Query(sql, groupId))
}

func processGames(sqlRows *sql.Rows, err error) ([]data_classes.GameData, error) {
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		mappedRows := sql_utils.ToSqlMap(sqlRows)

		games := make([]data_classes.GameData, len(mappedRows))

		for i, v := range mappedRows {
			u := data_classes.GameData{
				Id:             v["id"].(string),
				DateAdded:      v["date_added"].(time.Time),
				Description:    v["description"].(string),
				GroupId:        v["group_id"].(string),
				Name:           v["name"].(string),
				State:          v["state"].(int64),
				WinningBoardId: sql_utils.ObjToString("winning_board_id", v),
			}

			games[i] = u
		}

		return games, nil
	}
}
