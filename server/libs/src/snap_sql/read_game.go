package snap_sql

import (
	"data_classes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"sql_utils"
	"time"
	//enc "json_helpers"
)

// ---------------------- Read Group Functions ---------------------- //

/*
func HasGame(gameName string) (bool, error) {
	games, err := ReadGame(gameName)

	if err != nil {
		fmt.Println(err)
		return true, err // TODO: should this be true or false?
	} else if len(games) > 0 {
		return true, err
	} else {
		return false, err
	}
}
*/

func ReadGameFromId(gameId string) ([]data_classes.GameData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "readGameFromId.sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		rows, err := sql_utils.GetConnection().Query(string(sql), gameId)

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return processGames(rows, err)
		}
	}
}

func ReadGameFromName(gameName string) ([]data_classes.GameData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "readGameFromName.sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		rows, err := sql_utils.GetConnection().Query(string(sql), gameName)

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return processGames(rows, err)
		}
	}
}

func ReadGameInGroupFromName(groupId, gameName string) ([]data_classes.GameData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "readGameInGroupFromName.sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		rows, err := sql_utils.GetConnection().Query(string(sql), groupId, gameName)

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return processGames(rows, err)
		}
	}
}

func ReadAllGames(groupId string) ([]data_classes.GameData, error) {
	//sql := "SELECT * FROM _user WHERE email=$1"
	sql, err := ioutil.ReadFile(sql_utils.FilePath + "readAllGamesInGroup.sql")

	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {

		rows, err := sql_utils.GetConnection().Query(string(sql), groupId)

		if err != nil {
			fmt.Println(err)
			return nil, err
		} else {
			return processGames(rows, err)
		}
	}
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
