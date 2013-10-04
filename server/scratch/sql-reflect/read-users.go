package main

import (
	"data_classes"
	"database/sql"
	"fmt"
	"sql_utils"
	"sql_utils/caching"
	_ "time"
)

func ReadAllUsers() ([]*data_classes.UserProfile, error) {

	rows, err := sql_utils.GetConnection().Query(caching.CacheEntries.ReadAllUsers.Script)

	return processUserProfiles(rows, err)
}

func processUserProfiles(sqlRows *sql.Rows, err error) ([]*data_classes.UserProfile, error) {

	if err != nil {

		fmt.Println(err)
		return nil, err

	} else {

		mappedRows := sql_utils.ToSqlMap(sqlRows)

		profiles := loadFromMaps(mappedRows)

		return profiles, nil
	}
}

func loadFromMaps(dps []map[string]interface{}) []*data_classes.UserProfile {

	count := len(dps)
	users := make([]*data_classes.UserProfile, count)
	ptrs := make([]interface{}, count)

	for i := 0; i < count; i++ {
		users[i] = &data_classes.UserProfile{}
		ptrs[i] = users[i]
	}

	FromMaps(ptrs, dps)

	return users
}
