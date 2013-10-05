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

func HasGroup(groupName string) (bool, error) {
	groups, err := ReadGroup(groupName)

	if err != nil {
		fmt.Println(err)
		// TODO: should this be true or false?
		return false, err
	} else if len(groups) > 0 {
		return true, err
	} else {
		return false, err
	}
}

func ReadGroup(group_name string) ([]data_classes.GroupData, error) {

	sql := caching.CacheEntries.ReadGroup.Script

	rows, err := sql_utils.GetConnection().Query(sql, group_name)

	return processGroup(rows, err)
}

func processGroup(sqlRows *sql.Rows, err error) ([]data_classes.GroupData, error) {

	if err != nil {

		fmt.Println(err)
		return nil, err

	} else {

		mappedRows := sql_utils.ToSqlMap(sqlRows)

		groups := make([]data_classes.GroupData, len(mappedRows))

		for i, v := range mappedRows {
			u := data_classes.GroupData{
				Id:          v["id"].(string),
				GroupName:   v["group_name"].(string),
				Description: v["group_desc"].(string),
				DateAdded:   v["date_added"].(time.Time),
			}

			groups[i] = u
		}

		return groups, nil
	}
}
